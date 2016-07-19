'use strict';
import * as types from '../actions/action-types';

const initial_state = {
  active_ticker: null,
  selected_tab: 'overview'
};

const navigation = (state = initial_state, action = {}) => {
  switch (action.type) {
    case types.SET_ACTIVE_TICKER:
      return {
        ...state,
        active_ticker: action.payload.ticker
      }
    case types.SELECT_TAB:
      return {
        ...state,
        selected_tab: action.payload.tab
      }
    default:
      return state;
  }
}

export default navigation;
