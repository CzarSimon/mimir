export const kgKey = "AIzaSyCdlbj16sLBBpKc2Op0CgWCbAoOn91aPVs";


export const companyTerms =
  ['inc', 'corporation', 'plc', 'company', 'co.', 'inc.', 'com', 'limited', 'comp']

export const companyReplacements = {
  "corp": "Corporation"
}

export const devMode = !true;


export const baseUrl = (!devMode) ? '/api' : 'http://localhost:8000/api';
