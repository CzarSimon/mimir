'use strict'
import { createAction } from 'redux-actions';
import { kgSearch } from '../methods/google-api';
import { getRequest } from '../methods/api-methods';

/* --- Types --- */
export const RECIVE_COMPANY_DESC = 'mimir/description/RECIVE';
export const RECIVE_DESC_FAILURE = 'mimir/description/RECIVE_FAILURE';

/* --- Reducer --- */
const descriptions = (state = {}, action = {}) => {
  switch (action.type) {
    case RECIVE_COMPANY_DESC:
      const { ticker, description } = action.payload
      return {
        ...state,
        [ticker]: description
      };
    case RECIVE_DESC_FAILURE:
      const { error } = action.payload
      return {
        ...state,
        [action.payload.ticker]: error
      };
    default:
      return state;
  }
}
export default descriptions

/* --- Actions --- */
export const fetchCompanyDesc = (companyName, ticker) => (
  dispatch => (
    kgSearch(companyName)
    .then(desc => dispatch(reciveCompanyDesc(desc, ticker)))
    .catch(err => {
      console.log(err);
    })
  )
)

export const fetchDescription = (companyName, ticker) => (
  dispatch => (
    getRequest(`api/app/stock/description?ticker=${ticker}`)
    .then(res => dispatch(reciveCompanyDesc(res.description, ticker)))
    .catch(err => {
      console.log(err);
      return dispatch(fetchCompanyDesc(companyName, ticker));
    })
  )
)

export const reciveDescFailure = createAction(
  RECIVE_DESC_FAILURE, (error, ticker) => ({ error, ticker })
)

export const reciveCompanyDesc = createAction(
  RECIVE_COMPANY_DESC, (description, ticker) => ({ description, ticker })
)
