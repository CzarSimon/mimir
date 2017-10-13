import Auth0 from 'react-native-auth0';
import SHA256 from 'crypto-js/sha256'
import { AUTH_CONFIG } from '../credentials/config';
import * as store from './async-storage';
import { hash } from './helper-methods';

export const ID_TOKEN_KEY = 'mimir/user/token/id';
const ACCESS_TOKEN_KEY = 'mimir/user/token/access';
const TOKEN_EXPIRY_KEY = 'mimir/user/token/expiry';

const auth0 = new Auth0({
  domain: AUTH_CONFIG.domain,
  clientId: AUTH_CONFIG.clientId
});

export const login = provider => (
  auth0.webAuth.authorize({
    connection: provider,
    audience: `https://${AUTH_CONFIG.domain}/userinfo`,
    responseType: 'token id_token profile',
    scope: 'openid profile email'
  })
)

export const logout = () => {
  store.remove(ID_TOKEN_KEY);
  store.remove(ACCESS_TOKEN_KEY);
  store.remove(TOKEN_EXPIRY_KEY);
  store.remove(store.USER_ID_KEY);
}

export const storeCredentials = ({idToken, accessToken, expiresIn, id}) => {
  store.persist(ID_TOKEN_KEY, idToken);
  store.persist(ACCESS_TOKEN_KEY, accessToken);
  store.persist(TOKEN_EXPIRY_KEY, calcTokenExpiry(expiresIn));
  store.persist(store.USER_ID_KEY, id);
}

export const parseCredentials = accessToken => (
  auth0.auth
    .userInfo({token: accessToken})
    .then(profile => SHA256(profile.sub).toString())
)

const calcTokenExpiry = timestamp => (
  JSON.stringify((timestamp * 1000) + new Date().getTime())
)

export const getUserCredentials = () => (
  store.retrive(store.USER_ID_KEY)
  .then(id => {
    if (id === null) {
      throw new Error("No user id found");
    }
    return (
      store.retrive(ID_TOKEN_KEY)
      .then(token => {
        if (token === null) {
          throw new Error("No token found");
        }
        return ({id, token})
      })
    )
  })
)
