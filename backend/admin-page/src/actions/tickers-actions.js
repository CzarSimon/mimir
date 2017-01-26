import * as types from './action-types';
import { createAction } from 'redux-actions';
import KGSearch from 'google-kgsearch';
const getStock = require('get-stock');
import {
  parseCompanyDescription,
  parseCompanyName,
  parseImageUrl,
  createHttpObject
} from '../methods/helper-methods';
import { kgKey } from '../config';

export const reciveUntrackedTickers =
  createAction(types.RECIVE_UNTRACKED_TICKERS, tickers => (
    { tickers }
  ))

export const fetchUntrackedTickers = () => {
  return dispatch => {
    return fetch('http://localhost:8000/untracked-tickers')
    .then(res => res.json())
    .then(tickers => dispatch(reciveUntrackedTickers(tickers)))
    .catch(err => console.log("Error in fetch tickers: ", err))
  }
}

export const reciveTickerInfo =
  createAction(types.RECIVE_TICKER_INFO, (ticker, description, companyName, imageUrl) => {
    return {
      ticker,
      description,
      companyName,
      imageUrl
    }
  })

export const fetchTickerDescription = (ticker, name) => {
  return dispatch => {
    KGSearch(kgKey).search({query: name, limit: 1}, (err, res) => {
      const description = parseCompanyDescription(err, res)
      const imageUrl = parseImageUrl(err, res)
      dispatch(reciveTickerInfo(ticker, description, name, imageUrl))
    })
  }
}


export const fetchCompanyInfo = (ticker) => {
  return dispatch => {
    return getStock([ticker]).then(res => res.results.Name)
    .then(name => parseCompanyName(name))
    .then(name => {dispatch(fetchTickerDescription(ticker, name))})
    .catch(err => console.log("Error in fetch company info:", err))
  }
}


export const tickerTrackResponse = () =>
  createAction(types.START_TRACKING_TICKER, () => ({}))


export const startTrackingTicker = (ticker, name, description, token) => {
  token = "asfklaj"
  const httpObject = createHttpObject("POST", {ticker, name, description, token})
  console.log(httpObject.body);
  return dispatch => {
    return fetch('http://localhost:8000/track-ticker', httpObject)
    .then(res => {dispatch(tickerTrackResponse())})
  }
}
