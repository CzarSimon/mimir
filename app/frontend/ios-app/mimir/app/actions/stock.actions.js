import * as types from './action-types';
import { createAction } from 'redux-actions';
import { retrive_stock_data, retrive_historical_data } from './../methods/yahoo-api';
import { get_twitter_data } from '../methods/server/twitter-miner';

export const fetch_stock_data = (tickers) => {
  return (dispatch) => {
    return retrive_stock_data(tickers)
    .then(data => dispatch(recive_stock_data(data)))
    .catch(err => console.log("THERE WAS AN ERROR:", err))
  };
}

export const recive_stock_data = createAction(types.RECIVE_STOCK_DATA, data => (
  { data }
))

export const update_stock_data = (tickers) => {
  return (dispatch) => {
    return retrive_stock_data(tickers)
    .then(data => dispatch(recive_updated_stock_data(data)))
    .catch(err => console.log("THERE WAS AN ERROR:", err))
  }
}

export const recive_updated_stock_data = createAction(types.RECIVE_UPDATED_STOCK_DATA, data => (
  { data }
))

export const fetch_historical_data = (ticker) => {
  return (dispatch) => {
    return retrive_historical_data(ticker)
    .then(data => dispatch(recive_historical_data(data, ticker)))
    .catch(err => console.log("THERE WAS AN ERROR:", err))
  }
}

export const recive_historical_data = createAction(types.RECIVE_HISTORICAL_DATA, (data, ticker) => (
  {
    data,
    ticker
  }
))
