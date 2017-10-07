import auth0 from 'auth0-js';
import { browserHistory } from 'react-router';
import { AUTH_CONFIG } from '../credentials/config';
import * as store from './async-storage';
import { hash } from './helper-methods';

export const ID_TOKEN_KEY = 'mimir/user/token/id';
const ACCESS_TOKEN_KEY = 'mimir/user/token/access';
const TOKEN_EXPIRY_KEY = 'mimir/user/token/expiry';

export default class Auth {
  auth0 = new auth0.WebAuth({
    domain: AUTH_CONFIG.domain,
    clientID: AUTH_CONFIG.clientId,
    redirectUri: AUTH_CONFIG.callbackUrl,
    audience: `https://${AUTH_CONFIG.domain}/userinfo`,
    responseType: 'token id_token',
    scope: 'openid profile email'
  });

  constructor() {
    this.login = this.login;
    this.twitterLogin = this.twitterLogin;
    this.facebookLogin = this.facebookLogin;
    this.logout = this.logout;
    this.getProfile = this.getProfile;
    this.handleAuthentication = this.handleAuthentication;
    this.isAuthenticated = this.isAuthenticated
  }

  login = provider => {
    this.auth0.authorize({
      connection: provider
    });
  }

  facebookLogin = () => {
    this.login('facebook');
  }

  twitterLogin = () => {
    this.login('twitter');
  }

  logout = () => {
    store.remove(ID_TOKEN_KEY);
    store.remove(ACCESS_TOKEN_KEY);
    store.remove(TOKEN_EXPIRY_KEY);
    store.remove(store.USER_ID_KEY);
    browserHistory.push('/login');
  }

  isAuthenticated = () => {
    const userId = store.retrive(store.USER_ID_KEY);
    const token = store.retrive(ID_TOKEN_KEY);
    return (userId && token)
  }

  calcTokenExpiry = timestamp => (
    JSON.stringify((timestamp * 1000) + new Date().getTime())
  )

  persistUserId = userId => {
    store.persist(store.USER_ID_KEY, hash(userId));
  }

  setSession = (authResult, persistUser = true) => {
    const { idToken, idTokenPayload, accessToken, expiresIn } = authResult;
    store.persist(ID_TOKEN_KEY, idToken);
    store.persist(ACCESS_TOKEN_KEY, accessToken);
    if (persistUser) {
      this.persistUserId(idTokenPayload.sub);
    }
    store.persist(TOKEN_EXPIRY_KEY, this.calcTokenExpiry(expiresIn));
  }

  handleAuthentication = redirect => {
    this.auth0.parseHash((err, authResult) => {
      if (authResult && authResult.accessToken && authResult.idToken) {
        this.setSession(authResult);
        redirect(authResult.idToken);
      } else if (err) {
        console.log(err);
        alert(`Error: ${err.error}. Check the console for further details.`);
      }
    });
  }
}
