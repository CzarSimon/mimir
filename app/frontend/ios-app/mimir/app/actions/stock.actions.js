import * as types from './action-types';
import { createAction } from 'redux-actions';
import { retrive_stock_data } from './../methods/yahoo-api';
import { get_twitter_data } from '../methods/server/twitter-miner';

export const fetch_stock_data = (tickers) => {
  return (dispatch) => {
    return retrive_stock_data(tickers)
    .then(data => dispatch(recive_stock_data(data)))
  };
}

export const recive_stock_data = createAction(types.RECIVE_STOCK_DATA, data => (
  { data }
))

export const fetch_twitter_data = (tickers) => {
  return (dispatch) => {
    return get_twitter_data(tickers)
    .then(data => dispatch(recive_twitter_data(data)))
  }
}

export const recive_twitter_data = createAction(types.RECIVE_TWITTER_DATA, data => (
  { data }
))
