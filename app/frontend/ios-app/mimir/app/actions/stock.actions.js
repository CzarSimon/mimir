import * as types from './action-types';
import { createAction } from 'redux-actions';
import { retrive_stock_data } from './../methods/yahoo-api';
import { get_twitter_data } from '../methods/server/twitter-miner';

export const fetch_stock_data = (tickers) => {
  return (dispatch) => {
    return retrive_stock_data(tickers)
    .then(data => {
      return dispatch(recive_stock_data(data))
    })
  };
}

export const recive_stock_data = createAction(types.RECIVE_STOCK_DATA, data => (
  { data }
))

export const update_stock_data = (tickers) => {
  return (dispatch) => {
    return retrive_stock_data(tickers)
    .then(data => dispatch(recive_updated_stock_data(data)))
  }
}

export const recive_updated_stock_data = createAction(types.RECIVE_UPDATED_STOCK_DATA, data => (
  { data }
))
