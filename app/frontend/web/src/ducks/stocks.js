import { mapValues, omit } from 'lodash'
import { createAction } from 'redux-actions'
import { retriveStockData, retriveHistoricalData } from './../methods/yahoo-api'

/* --- Types --- */
export const RECIVE_STOCK_DATA = 'mimir/stockData/RECIVE';
export const FETCH_STOCK_DATA = 'mimir/stockData/FETCH'; //Add failure scenario
export const RECIVE_UPDATED_STOCK_DATA = 'mimir/stockData/RECIVE_UPDATE';
export const UPDATE_STOCK_DATA = 'mimir/stockData/UPDATE'; //Add failure scenario
export const RECIVE_HISTORICAL_DATA = 'mimir/stockData/historical/RECIVE';
export const FETCH_HISTORICAL_DATA = 'mimir/stockData/historical/FETCH'; //Add failure scenario
export const DELETE_STOCK_DATA = 'mimir/stockData/DELETE';

/* --- Reducer --- */
const initialState = {
  data: [],
  loaded: false
};

const stocks = (state = initialState, action = {}) => {
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
    case DELETE_STOCK_DATA:
      return {
        ...state,
        data: omit(state.data, action.payload.ticker)
      }
    default:
      return state;
  }
}
export default stocks

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

export const deleteStockData = createAction(
  DELETE_STOCK_DATA, ticker => ({ ticker })
)
