import * as types from '../actions/action-types';
import { filter } from 'lodash';

const initalState = {
  data: [],
  loaded: false
}

const tickers = (state = initalState, action = {}) => {
  switch (action.type) {
    case types.RECIVE_UNTRACKED_TICKERS:
      return {
        ...state,
        data: action.payload.tickers,
        loaded: true
      };
    case types.START_TRACKING_TICKER:
      return {
        ...state,
        data: filter(state.data, ticker => (
          ticker.tickerName !== action.payload.ticker
        ))
      };
    default:
      return state;
  }
}

export default tickers;
