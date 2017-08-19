'use strict'
import { createAction } from 'redux-actions';
import { toURL } from '../methods/helper-methods';
import { getRequest, createTickerQuery } from '../methods/api-methods';

/* --- Types --- */
export const FETCH_SEARCH_RESULTS = 'mimir/searchResults/FETCH';
export const RECIVE_SEARCH_RESULTS = 'mimir/searchResults/RECIVE';
export const FETCH_SEARCH_SUGESTIONS = 'mimir/searchSugestions/FETCH';
export const RECIVE_SEARCH_SUGESTIONS = 'mimir/searchSugestions/RECIVE';
export const SEARCH_KEYBOARD_UP = 'mimir/search/keyboard/UP';
export const SEARCH_KEYBOARD_DOWN = 'mimir/search/keyboard/DOWN';
export const UPDATE_QUERY = 'mimir/query/UPDATE';

const initialState = {
  query: null,
  keyboardDown: true,
  results: [],
  sugestions: []
}

/* --- Reducer --- */
const search = (state = initialState, action = {}) => {
  switch (action.type) {
    case RECIVE_SEARCH_RESULTS:
      return {
        ...state,
        results: action.payload.results
      };
    case SEARCH_KEYBOARD_UP:
      return {
        ...state,
        keyboardDown: false
      }
    case SEARCH_KEYBOARD_DOWN:
      return {
        ...state,
        keyboardDown: true
      }
    case UPDATE_QUERY:
      return {
        ...state,
        query: action.payload.query
      }
    case RECIVE_SEARCH_SUGESTIONS:
      return {
        ...state,
        sugestions: action.payload.sugestions
      }
    default:
      return state;
  }
}
export default search

/* --- Actions --- */
export const cancelSearch = createAction(SEARCH_KEYBOARD_DOWN);

export const activateSearchKeyboard = createAction(SEARCH_KEYBOARD_UP);

export const reciveSearchResults = createAction(
  RECIVE_SEARCH_RESULTS, results => ({ results })
)

export const fetchSearchResults = query => {
  return dispatch => (
    getRequest(`api/search?query=${query}`)
    .then(res => dispatch(reciveSearchResults(res)))
    .catch(err => {
      console.log(err);
    })
  );
}

export const reciveSearchSugestions = createAction(
  RECIVE_SEARCH_SUGESTIONS, sugestions => ({ sugestions })
);

export const fetchSearchSugestions = tickers => {
  return dispatch => (
    getRequest('api/search/sugestions' + createTickerQuery(tickers))
    .then(res => dispatch(reciveSearchSugestions(res)))
    .catch(err => {
      console.log(err);
    })
  );
}

export const updateQuery = createAction(UPDATE_QUERY, query => ({ query }))

export const updateAndRunQuery = query => (
  dispatch => (query.length > 0) ? (
    Promise.all([
      dispatch(fetchSearchResults(query)),
      dispatch(updateQuery(query))
    ])
  ) : dispatch(updateQuery(query))
);
