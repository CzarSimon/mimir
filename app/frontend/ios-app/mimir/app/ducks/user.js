'use strict'
import { createAction } from 'redux-actions'
import twitterData from './twitter-data';
import { without, concat } from 'lodash';
import { retriveObject} from '../methods/async-storage';

/* --- Types --- */
export const RECIVE_USER = 'RECIVE_USER'
export const FETCH_USER = 'FETCH_USER'
export const CREATE_NEW_USER = 'CREATE_NEW_USER'
export const ADD_TICKER = 'ADD_TICKER'
export const REMOVE_TICKER = 'REMOVE_TICKER'
export const TOGGLE_MODIFIABLE = 'TOGGLE_MODIFIABLE'
export const CLEAR_SEARCH_HISTORY = 'CLEAR_SEARCH_HISTORY'

const initialState = {
  id: null,
  name: null,
  tickers: [],
  searchHistory: ['twitter', 'GOOG', 'Nvid'],
  twitterData: twitterData(),
  modifiable: false,
  loaded: false
};

// modifiable is chaned to true on recive user --> modifiable is stored in async-storage

/* --- Reducer --- */
export default user = (state = initialState, action = {}) => {
  switch (action.type) {
    case RECIVE_USER:
      return {
        ...state,
        ...action.payload.user,
        loaded: true
      };
    case CREATE_NEW_USER:
      return {
        ...state,
        ...action.payload.user,
        loaded: true
      };
    case ADD_TICKER:
      const { ticker } = action.payload
      return {
        ...state,
        tickers: (
          (!state.tickers.includes(ticker))
          ? concat(state.tickers, ticker)
          : state.tickers
        )
      };
    case REMOVE_TICKER:
      return {
        ...state,
        tickers: without(state.tickers, action.payload.ticker)
      };
    case RECIVE_TWITTER_DATA:
      return {
        ...state,
        twitter_data: twitterData(state, action)
      };
    case TOGGLE_MODIFIABLE:
      return {
        ...state,
        modifiable: !state.modifiable
      }
    case CLEAR_SEARCH_HISTORY:
      return {
        ...state,
        searchHistory: []
      }
    default:
      return state;
  }
}

/* --- Actions --- */
export const fetchUser = () => {
  return dispatch => (
    retriveObject("user")
    .then(user => dispatch(recive_user(user)))
  )
}

export const reciveUser = createAction(RECIVE_USER, user => ({ user }))

export const createNewUser = createAction(CREATE_NEW_USER, user => ({ user }))

export const add_ticker = createAction(ADD_TICKER, ticker => ({ ticker }))

export const removeRicker = createAction(REMOVE_TICKER, ticker => ({ ticker }))

export const toggleModifiable = createAction(TOGGLE_MODIFIABLE)

export const clearSearchHistory = createAction(CLEAR_SEARCH_HISTORY)
