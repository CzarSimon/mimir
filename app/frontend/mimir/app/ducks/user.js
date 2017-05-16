'use strict'
import { createAction } from 'redux-actions'
import twitterData from './twitter-data'
import { RECIVE_TWITTER_DATA } from './twitter-data'
import { without, concat, uniq } from 'lodash'
import { retriveObject } from '../methods/async-storage'

/* --- Types --- */
export const RECIVE_USER = 'RECIVE_USER'
export const FETCH_USER = 'FETCH_USER'
export const CREATE_NEW_USER = 'CREATE_NEW_USER'
export const ADD_TICKER = 'ADD_TICKER'
export const REMOVE_TICKER = 'REMOVE_TICKER'
export const CLEAR_SEARCH_HISTORY = 'CLEAR_SEARCH_HISTORY'
export const ADD_TO_SEARCH_HISTORY = 'ADD_TO_SEARCH_HISTORY'

const initialState = {
  id: null,
  name: null,
  tickers: [],
  searchHistory: ['twitter', 'GOOG', 'Nvd'],
  twitterData: twitterData(),
  modifiable: false,
  loaded: false
}

// modifiable is chaned to true on recive user --> modifiable is stored in async-storage

/* --- Reducer --- */
const user = (state = initialState, action = {}) => {
  switch (action.type) {
    case RECIVE_USER:
      return {
        ...state,
        ...action.payload.user,
        loaded: true
      }
    case CREATE_NEW_USER:
      return {
        ...state,
        ...action.payload.user,
        loaded: true
      }
    case ADD_TICKER:
      const { ticker } = action.payload
      return {
        ...state,
        tickers: (
          (!state.tickers.includes(ticker))
          ? concat(state.tickers, ticker)
          : state.tickers
        )
      }
    case REMOVE_TICKER:
      return {
        ...state,
        tickers: without(state.tickers, action.payload.ticker)
      }
    case RECIVE_TWITTER_DATA:
      return {
        ...state,
        twitterData: twitterData(state, action)
      }
    case CLEAR_SEARCH_HISTORY:
      return {
        ...state,
        searchHistory: []
      }
    case ADD_TO_SEARCH_HISTORY:
      return {
        ...state,
        searchHistory: uniq(concat(action.payload.query, state.searchHistory))
      }
    default:
      return state
  }
}
export default user

/* --- Actions --- */
export const fetchUser = () => {
  return dispatch => (
    retriveObject("user")
    .then(user => dispatch(reciveUser(user)))
  )
}

export const reciveUser = createAction(RECIVE_USER, user => ({ user }))

export const createNewUser = createAction(CREATE_NEW_USER, user => ({ user }))

export const addTicker = createAction(ADD_TICKER, ticker => ({ ticker }))

export const removeTicker = createAction(REMOVE_TICKER, ticker => ({ ticker }))

export const clearSearchHistory = createAction(CLEAR_SEARCH_HISTORY)

export const addToSearchHistory = createAction(
  ADD_TO_SEARCH_HISTORY, query => ({ query })
)

export const updateUserWithTicker = ticker => {

}
