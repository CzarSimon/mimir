import * as types from './action-types';
import { createAction } from 'redux-actions';
import { getNameYahoo, getCompanyInfo } from '../methods/stock-info';
import { createHttpObject} from '../methods/helper-methods';
import { baseUrl } from '../config';

export const reciveUntrackedTickers =
  createAction(types.RECIVE_UNTRACKED_TICKERS, tickers => (
    { tickers }
  ))

export const fetchUntrackedTickers = token => {
  const httpObject = createHttpObject('GET', token)
  return dispatch => {
    return fetch(`${baseUrl}/untracked-tickers`, httpObject)
    .then(res => res.json())
    .then(tickers => dispatch(reciveUntrackedTickers(tickers)))
    .catch(err => console.log("Error in fetch tickers: ", err))
  }
}

export const reciveTickerInfo =
  createAction(types.RECIVE_TICKER_INFO,
    (ticker, description, companyName, imageUrl, website) => {
      return {
        ticker,
        description,
        companyName,
        imageUrl,
        website
      }
  })


export const fetchTickerInfo = (ticker, name) => dispatch => getCompanyInfo(name, ticker, dispatch)


export const fetchCompanyInfo = (ticker) => {
  return dispatch => {
    return getNameYahoo(ticker)
    .then(name => { dispatch(fetchTickerInfo(ticker, name)) })
    .catch(err => console.log("Full name fecth failure", err))
  }
}


export const tickerTrackResponse = () =>
  createAction(types.START_TRACKING_TICKER, () => ({}))


export const startTrackingTicker =
(ticker, name, description, imageUrl, website, token) => {
  const body = {ticker, name, description, imageUrl, website, token}
  const httpObject = createHttpObject("POST", token, body)
  return dispatch => {
    return fetch(`${baseUrl}/track-ticker`, httpObject)
    .then(res => res.json())
    .then(json => {
      alert(json.Response)
      dispatch(tickerTrackResponse())
    })
  }
}
