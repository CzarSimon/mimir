import * as types from '../actions/action-types';
import twitter_data from './twitter_data.reducer';
import { without, concat } from 'lodash';

const initial_state = {
  id: null,
  name: null,
  tickers: [],
  twitter_data: twitter_data(),
  loaded: false
};

const user = (state = initial_state, action = {}) => {
  switch (action.type) {
    case types.RECIVE_USER:
      return {
        ...state,
        ...action.payload
      };
    case types.ADD_TICKER:
      return {
        ...state,
        tickers: (
          (!state.tickers.includes(action.payload.ticker))
          ? concat(state.tickers, action.payload.ticker)
          : state.tickers
        )
      };
    case types.REMOVE_TICKER:
      return {
        ...state,
        tickers: without(state.tickers, action.payload.ticker)
      };
    case types.RECIVE_TWITTER_DATA:
      return {
        ...state,
        twitter_data: twitter_data(state, action)
      };
    default:
      return state;
  }
}

export default user;
