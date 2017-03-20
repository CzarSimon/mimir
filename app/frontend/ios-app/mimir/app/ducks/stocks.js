'use strict'
import { mapValues } from 'lodash'
import { createAction } from 'redux-actions'
import { retriveStockData, retriveHistoricalData } from './../methods/yahoo-api';
import { getTwitterData } from '../methods/server/twitter-miner';

/* --- Types --- */
export const RECIVE_STOCK_DATA = 'RECIVE_STOCK_DATA'
export const FETCH_STOCK_DATA = 'FETCH_STOCK_DATA' //Add failure scenario
export const RECIVE_UPDATED_STOCK_DATA = 'RECIVE_UPDATED_STOCK_DATA'
export const UPDATE_STOCK_DATA = 'UPDATE_STOCK_DATA' //Add failure scenario
export const RECIVE_HISTORICAL_DATA = 'RECIVE_HISTORICAL_DATA'
export const FETCH_HISTORICAL_DATA = 'FETCH_HISTORICAL_DATA' //Add failure scenario

/* --- Reducer --- */
const initialState = {
  data: [],
  loaded: false
};

export default stocks = (state = initialState, action = {}) => {
  switch (action.type) {
    case RECIVE_STOCK_DATA:
      return {
        ...state,
        data: {
          ...state.data,
          ...action.payload.data
        },
        loaded: true
      };
    case RECIVE_UPDATED_STOCK_DATA:
      return {
        ...state,
        data: mapValues(state.data, (val) => (
          {
            ...val,
            ...action.payload.data[val.Symbol]
          }
        )),
        loaded: true
      };
    case RECIVE_HISTORICAL_DATA:
      const { ticker, data } = action.payload
      return {
        ...state,
        data: {
          ...state.data,
          [ticker]: {
            ...state.data[ticker],
            historicalData: data
          }
        },
        loaded: true
      }
    default:
      return state;
  }
}

/* --- Actions --- */
export const fetchStockData = tickers => (
  dispatch => (
    retriveStockData(tickers)
    .then(data => dispatch(reciveStockData(data)))
    .catch(err => console.log("THERE WAS AN ERROR:", err))
  )
)

export const reciveStockData = createAction(
  RECIVE_STOCK_DATA, data => ({ data })
)

export const updateStockData = tickers => (
  dispatch => (
    retriveStockData(tickers)
    .then(data => dispatch(reciveUpdatedStockData(data)))
    .catch(err => console.log("THERE WAS AN ERROR:", err))
  )
)

export const reciveUpdatedStockData = createAction(
  RECIVE_UPDATED_STOCK_DATA, data => ({ data })
)

export const fetchHistoricalData = ticker => (
  dispatch => (
    retriveHistoricalData(ticker)
    .then(data => dispatch(reciveHistoricalData(data, ticker)))
    .catch(err => console.log("THERE WAS AN ERROR:", err))
  )
)

export const reciveHistoricalData = createAction(
  RECIVE_HISTORICAL_DATA, (data, ticker) => ({ data, ticker })
)
