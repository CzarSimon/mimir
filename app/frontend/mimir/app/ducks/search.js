'use strict'
import { createAction } from 'redux-actions';
import { toURL } from '../methods/helper-methods';
import { getRequest } from '../methods/api-methods';

/* --- Types --- */
export const FETCH_SEARCH_RESULTS = 'mimir/searchResults/FETCH';
export const RECIVE_SEARCH_RESULTS = 'mimir/searchResults/RECIVE';
export const TOGGLE_KEYBOARD_UP = 'mimir/keyboard/TOGGLE';
export const UPDATE_QUERY = 'mimir/query/UPDATE';

const initialState = {
  query: null,
  keyboardUp: false,
  results: []
}

/* --- Reducer --- */
const search = (state = initialState, action = {}) => {
  switch (action.type) {
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
export const toggleKeyboardUp = createAction(TOGGLE_KEYBOARD_UP);

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

export const updateQuery = createAction(UPDATE_QUERY, query => ({ query }))

export const updateAndRunQuery = query => (
  dispatch => (query.length > 0) ? (
    Promise.all([
      dispatch(fetchSearchResults(query)),
      dispatch(updateQuery(query))
    ])
  ) : dispatch(updateQuery(query))
);
