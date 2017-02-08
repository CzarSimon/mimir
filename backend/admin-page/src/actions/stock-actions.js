import * as types from './action-types';
import { createAction } from 'redux-actions';
import {
  createHttpObject,
  createPath,
  objectArrayToObject
} from '../methods/helper-methods';


export const reciveTrackedStocks =
  createAction(types.RECIVE_TRACKED_STOCKS, stocks => ({stocks}))


export const fetchTrackedStocks = token => {
  const httpObject = createHttpObject('GET', token)
  return dispatch => {
    return fetch(createPath('/tracked-stocks'), httpObject)
    .then(res => res.json())
    .then(res => {
      const stockData = objectArrayToObject(res, 'Ticker')
      dispatch(reciveTrackedStocks(stockData))
    })
  }
}

export const updateFilter = createAction(types.UPDATE_FILTER, filterTerm => ({filterTerm}))
