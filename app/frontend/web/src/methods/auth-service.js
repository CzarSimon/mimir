import auth0 from 'auth0-js';
import { browserHistory } from 'react-router';
import { AUTH_CONFIG } from '../credentials/config';
import * as store from './async-storage';
import { hash } from './helper-methods';
import { color } from '../styles/main';

export const ID_TOKEN_KEY = 'mimir/user/token/id';
const ACCESS_TOKEN_KEY = 'mimir/user/token/access';

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
    this.handleAuthentication = this.handleAuthentication.bind(this);
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
    store.remove(store.USER_ID_KEY);
    browserHistory.push('/login');
  }

  isAuthenticated = () => {
    const userId = store.retrive(store.USER_ID_KEY);
    const token = store.retrive(ID_TOKEN_KEY);
    return (userId && token)
  }

  setSession = (redirect, authResult) => {
    console.log(authResult);
    store.persist(ID_TOKEN_KEY, authResult.idToken);
    store.persist(ACCESS_TOKEN_KEY, authResult.accessToken);
    const userId = hash(authResult.idTokenPayload.sub);
    store.persist(store.USER_ID_KEY, userId);
    redirect(authResult.idToken);
  }

  handleAuthentication = redirect => {
    this.auth0.parseHash((err, authResult) => {
      if (authResult && authResult.accessToken && authResult.idToken) {
        this.setSession(redirect, authResult);
      } else if (err) {
        console.log(err);
        alert(`Error: ${err.error}. Check the console for further details.`);
      }
    });
  }
}
