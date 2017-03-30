'use strict'
import { createAction } from 'redux-actions';
import { kgSearch } from '../methods/google-api';

/* --- Types --- */
export const RECIVE_COMPANY_DESC = 'RECIVE_COMPANY_DESC'
export const RECIVE_DESC_FAILURE = 'RECIVE_DESC_FAILURE'

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
    .catch(err => dispatch(reciveDescFailure(err, ticker)))
  )
)

export const reciveDescFailure = createAction(
  RECIVE_DESC_FAILURE, (error, ticker) => ({ error, ticker })
)

export const reciveCompanyDesc = createAction(
  RECIVE_COMPANY_DESC, (description, ticker) => ({ description, ticker })
)
