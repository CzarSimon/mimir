import * as names from './route-names';

export const MAIN_ROUTE = {name: names.MAIN, title: 'mimir', index: 0};
export const companyPageRoute = title => (
  {
    name: names.COMPANY_PAGE,
    index: 1, title
  }
)
