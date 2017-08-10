'use strict'
import { createAction } from 'redux-actions';
import twitterData from './twitter-data';
import { RECIVE_TWITTER_DATA } from './twitter-data';
import { without, concat, uniq } from 'lodash';
import { retrive, persist, USER_ID_KEY } from '../methods/async-storage';
import { toURL } from '../methods/helper-methods';

/* --- Types --- */
export const RECIVE_USER = 'mimir/user/RECIVE';
export const FETCH_USER = 'mimir/user/FETCH';
export const ADD_TICKER = 'mimir/ticker/ADD';
export const DELETE_TICKER = 'mimir/ticker/DELETE';
export const CLEAR_SEARCH_HISTORY = 'mimir/searchHistory/CLEAR';
export const APPEND_SEARCH_HISTORY = 'mimir/searchHistory/APPEND';

const initialState = {
  id: null,
  email: null,
  joinDate: null,
  tickers: [],
  searchHistory: [],
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
      };
    case ADD_TICKER:
      const { ticker } = action.payload;
      return {
        ...state,
        tickers: (
          (!state.tickers.includes(ticker))
          ? concat(state.tickers, ticker)
          : state.tickers
        )
      };
    case DELETE_TICKER:
      return {
        ...state,
        tickers: without(state.tickers, action.payload.ticker)
      };
    case CLEAR_SEARCH_HISTORY:
      return {
        ...state,
        searchHistory: []
      };
    case APPEND_SEARCH_HISTORY:
      return {
        ...state,
        searchHistory: uniq(concat(action.payload.query, state.searchHistory))
      };
    default:
      return state;
  }
}

export default user

/* --- Actions --- */
export const getUser = () => {
  return dispatch => (
    retrive(USER_ID_KEY)
    .then(userId => {
      if (userId) {
        return dispatch(fetchUser(userId));
      } else {
        return dispatch(fetchNewUser());
      }
    })
  )
}

const fetchUser = userId => {
  return dispatch => (
    fetch(toURL('api/app/user?id=' + userId))
    .then(res => res.json())
    .then(user => dispatch(reciveUser(user)))
    .catch(err => {
      console.log(err);
    })
  );
}

const fetchNewUser = () => {
  return dispatch => (
    fetch(toURL('api/app/user'), { method: 'POST' })
    .then(res => res.json())
    .then(user => {
      persist(USER_ID_KEY, user.id)
      return dispatch(reciveUser(user))
    })
    .catch(err => {
      console.log(err);
    })
  );
}

export const reciveUser = createAction(RECIVE_USER, user => ({ user }));

export const addTicker = createAction(ADD_TICKER, ticker => ({ ticker }));

export const removeTicker = createAction(DELETE_TICKER, ticker => ({ ticker }));

export const clearSearchHistory = createAction(CLEAR_SEARCH_HISTORY);

export const addToSearchHistory = createAction(
  APPEND_SEARCH_HISTORY, query => ({ query })
);
