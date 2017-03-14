'use strict'
import { createAction } from 'redux-actions';

/* --- Types --- */
export const SET_ACTIVE_TICKER = 'SET_ACTIVE_TICKER'
export const SELECT_TAB = 'SELECT_TAB'

const initialState = {
  activeTicker: null,
  selectedTab: 'news',
  articleUrl: null
};

/* --- Reducer --- */
export default navigation = (state = initialState, action = {}) => {
  switch (action.type) {
    case SET_ACTIVE_TICKER:
      return {
        ...state,
        activeTicker: action.payload.ticker
      }
    case SELECT_TAB:
      return {
        ...state,
        selectedTab: action.payload.tab
      }
    default:
      return state;
  }
}

/* --- Actions --- */
export const setActiveTicker = createAction(SET_ACTIVE_TICKER, ticker => ({ ticker }))

export const selectTab = createAction(SELECT_TAB, tab => ({ tab }))
