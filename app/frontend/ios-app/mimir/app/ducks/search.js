'use strict'
import { createAction } from 'redux-actions';
import socket from '../methods/server/socket';

/* --- Types --- */
export const TOGGLE_SEARCH_ACTIVE = 'TOGGLE_SEARCH_ACTIVE'
export const FETCH_SEARCH_RESULTS = 'FETCH_SEARCH_RESULTS'
export const RECIVE_SEARCH_RESULTS = 'RECIVE_SEARCH_RESULTS'
export const TOGGLE_KEYBOARD_UP = 'TOGGLE_KEYBOARD_UP'
export const UPDATE_QUERY = 'UPDATE_QUERY'

const initialState = {
  active: false,
  query: null,
  placeholder: 'Search tickers...',
  keyboardUp: false,
  results: []
}

/* --- Reducer --- */
const search = (state = initialState, action = {}) => {
  switch (action.type) {
    case TOGGLE_SEARCH_ACTIVE:
      return {
        ...state,
        active: !state.active
      };
    case RECIVE_SEARCH_RESULTS:
      return {
        ...state,
        results: action.payload.results
      };
    case TOGGLE_KEYBOARD_UP:
      return {
        ...state,
        keyboardUp: !state.keyboardUp
      }
    case UPDATE_QUERY:
      return {
        ...state,
        query: action.payload.query
      }
    default:
      return state;
  }
}
export default search

/* --- Actions --- */
export const toggleSearchActive = createAction(TOGGLE_SEARCH_ACTIVE)

export const toggleKeyboardUp = createAction(TOGGLE_KEYBOARD_UP)

export const reciveSearchResults = createAction(
  RECIVE_SEARCH_RESULTS, results => ({ results })
)

export const fetchSearchResults = query => (() =>
  socket.emit(FETCH_SEARCH_RESULTS, { query })
)

export const updateQuery = query => {
  fetchSearchResults(query)
  return createAction(UPDATE_QUERY, query => ({ query }))
}
