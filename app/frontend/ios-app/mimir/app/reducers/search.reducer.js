'use strict';
import * as types from '../actions/action-types';

const initial_state = {
  active: false,
  query: null,
  placeholder: "Search tickers...",
  results: []
}

const search = (state = initial_state, action = {}) => {
  switch (action.type) {
    case types.TOGGLE_SEARCH_ACTIVE:
      return {
        ...state,
        active: !state.active
      };
    case types.RECIVE_SEARCH_RESULTS:
      return {
        ...state,
        results: action.payload.results
      };
    default:
      return state;
  }
}

export default search;
