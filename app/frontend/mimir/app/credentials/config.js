export const DEV_MODE = !false;
export const SERVER_URL = (DEV_MODE) ? 'http://localhost:8080/' : 'https://mimir.news/';
export const KG_KEY ='AIzaSyCdlbj16sLBBpKc2Op0CgWCbAoOn91aPVs';

const callbackUrl = (!DEV_MODE) ?
    'https://mimir.news/callback' :
    'http://localhost:3001/callback';

export const AUTH_CONFIG = {
  domain: 'mimir.eu.auth0.com',
  clientId: 'Uh3iD5AxoQYzrGjSeATWe6RCm7LBhPtM',
  callbackUrl
}
