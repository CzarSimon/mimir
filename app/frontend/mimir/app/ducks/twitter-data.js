'use strict'
import { createAction } from 'redux-actions';
import { toURL } from '../methods/helper-methods';
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
        data: _.keyBy(action.payload.data, 'ticker'),
        loaded: true
      }
    default:
      return state;
  }
}
export default twitterData

/* --- Actions --- */
export const fetchTwitterData = tickers => {
  const endpoint = toURL('api/app/twitter-data' + createTickerQuery(tickers));
  return dispatch => (
    fetch(endpoint)
    .then(res => res.json())
    .then(twitterData => dispatch(reciveTwitterData(twitterData)))
    .catch(err => {
      console.log(err);
    })
  );
}

export const reciveTwitterData = createAction(RECIVE_TWITTER_DATA, data => ({ data }))

const createTickerQuery = tickers => (
  '?' + _.join(_.map(tickers, ticker => 'ticker=' + ticker), '&')
);
