'use strict'
import { getUser } from './user';
import { fetchTwitterData } from './twitter-data';
import { fetchStockData } from './stocks';

/* --- Types --- */
export const USER_LOGON = 'mimir/user/LOGON';

/* --- Reducer --- */

/* --- Actions --- */
export const logonUser = () => {
  return (dispatch, getState) => {
    return dispatch(getUser()).then(() => {
      const { tickers } = getState().user;
      return Promise.all([
        dispatch(fetchTwitterData(tickers)),
        dispatch(fetchStockData(tickers))
      ])
    })
  }
}
