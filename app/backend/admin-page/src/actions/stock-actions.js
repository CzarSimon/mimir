import * as types from './action-types';
import { createAction } from 'redux-actions';
import {
  createHttpObject,
  createPath,
  objectArrayToObject
} from '../methods/helper-methods';


export const reciveTrackedStocks = createAction(
  types.RECIVE_TRACKED_STOCKS, stocks => ({stocks})
)


export const fetchTrackedStocks = token => {
  const httpObject = createHttpObject('GET', token)
  return dispatch => {
    return fetch(createPath('/tracked-stocks'), httpObject)
    .then(res => res.json()).then(res => {
      const stockData = objectArrayToObject(res, 'Ticker')
      dispatch(reciveTrackedStocks(stockData))
    })
  }
}


export const updateFilter = createAction(
  types.UPDATE_FILTER, filterTerm => ({ filterTerm })
)


export const removeUntrackedStock = createAction(
  types.REMOVE_UNTACKED_STOCK, ticker => ({ ticker })
)


export const untrackStock = (ticker, token) => {
  const httpObject = createHttpObject('POST', token, { ticker })
  return dispatch => {
    return fetch(createPath('/untrack-stock'), httpObject)
    .then(res => res.json()).then(res => {
      dispatch(removeUntrackedStock(ticker))
    })
    .catch(err => console.log(err))
  }
}


export const toggleEditMode = createAction(
  types.TOGGLE_EDIT_MODE, ticker => ({ ticker })
)


export const saveStockInfo = createAction(
  types.SAVE_STOCK_INFO, (ticker, description) => ({ ticker, description })
)


export const updateStockInfo = (ticker, description, token) => {
  console.log(token, {ticker, description});
  const httpObject = createHttpObject('POST', token, { ticker, description })
  return dispatch => {
    return (
      fetch(createPath('/update-stock-info'), httpObject)
      .then(res => res.json()).then(res => {
        dispatch(saveStockInfo(ticker, description))
      })
    )
  }
}
