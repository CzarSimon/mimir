'use strict'
import { getUser, fetchUser } from './user';
import { fetchTwitterData } from './twitter-data';
import { fetchStockData } from './stocks';
import { postRequest } from '../methods/api-methods';

/* --- Reducer --- */

/* --- Actions --- */
export const logonUser = () => {
  return (dispatch, getState) => {
    return dispatch(getUser()).then(() => {
      const { tickers, id } = getState().user;
      recordSession(id);
      return Promise.all([
        dispatch(fetchTwitterData(tickers)),
        dispatch(fetchStockData(tickers))
      ])
    })
    .catch(console.log)
  }
}

export const fetchAndInitalizeUser = userId => (
  (dispatch, getState) => {
    return dispatch(fetchUser(userId)).then(() => {
      const { tickers } = getState().user;
      recordSession(userId);
      return Promise.all([
        dispatch(fetchTwitterData(tickers)),
        dispatch(fetchStockData(tickers))
      ])
    })
  }
)

const recordSession = userId => {
  postRequest('api/app/user/session', { id: userId })
  .catch(err => {
    console.log(err);
  });
}
