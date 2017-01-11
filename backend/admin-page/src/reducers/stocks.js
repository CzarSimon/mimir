import * as types from '../actions/action-types';

const initalState = {
  data: [],
  loaded: false
}

const stocks = (state = initalState, action = {}) => {
  switch (action.type) {
    case types.RECIVE_TRACKED_STOCKS:
      return {
        ...state,
        data: action.payload.stocks,
        loaded: true
      };
    case types.SAVE_STOCK_INFO:
      return state;
    default:
      return state;
  }
}

export default stocks;
