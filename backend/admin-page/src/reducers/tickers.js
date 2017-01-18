import * as types from '../actions/action-types';
import { filter } from 'lodash';
import { objectArrayToObject } from '../methods/helper-methods';

const initalState = {
  data: {},
  loaded: false
}

const tickers = (state = initalState, action = {}) => {
  switch (action.type) {
    case types.RECIVE_UNTRACKED_TICKERS:
      return {
        ...state,
        data: {
          ...state.data,
          ...objectArrayToObject(action.payload.tickers, 'Name')
        },
        loaded: true
      };
    case types.RECIVE_TICKER_DESCRIPTION:
      return {
        ...state,
        data: {
          ...state.data,
          [action.payload.ticker]: {
            ...state.data[action.payload.ticker],
            description: action.payload.description
          }
        }
      }
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
