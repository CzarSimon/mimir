import * as types from '../actions/action-types';
import twitter_data from './twitter_data.reducer';
import { without, concat } from 'lodash';

const initial_state = {
  id: null,
  name: null,
  tickers: [],
  twitter_data: twitter_data(),
  modifiable: false,
  loaded: false
};

// modifiable is chaned to true on recive user --> modifiable is stored in async-storage

const user = (state = initial_state, action = {}) => {
  switch (action.type) {
    case types.RECIVE_USER:
      return {
        ...state,
        ...action.payload.user,
        loaded: true
      };
    case types.CREATE_NEW_USER:
      return {
        ...state,
        ...action.payload.user,
        loaded: true
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
    case types.TOGGLE_MODIFIABLE:
      return {
        ...state,
        modifiable: !state.modifiable
      }
    default:
      return state;
  }
}

export default user;
