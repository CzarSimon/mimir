'use strict'
import { createAction } from 'redux-actions';
import { getRequest, createTickerQuery } from '../methods/api-methods';
import _ from 'lodash';

/* --- Types --- */
export const RECIVE_TWITTER_DATA = 'mimir/twitterData/RECIVE';
export const FETCH_TWITTER_DATA = 'mimir/twitterData/FETCH'; //Add failure scenario

const initialState = {
  data: {},
  loaded: false
}

/* --- Reducer --- */
const twitterData = (state = initialState, action = {}) => {
  switch (action.type) {
    case RECIVE_TWITTER_DATA:
      return {
        data: {
          ...state.data,
          ..._.keyBy(action.payload.data, 'ticker'),
        },
        loaded: true
      }
    default:
      return state;
  }
}
export default twitterData

/* --- Actions --- */
export const fetchTwitterData = tickers => {
  const route = 'api/app/twitter-data' + createTickerQuery(tickers);
  return dispatch => (
    getRequest(route)
    .then(twitterData => dispatch(reciveTwitterData(twitterData)))
    .catch(err => {
      console.log(err);
    })
  );
}

export const reciveTwitterData = createAction(RECIVE_TWITTER_DATA, data => ({ data }))
