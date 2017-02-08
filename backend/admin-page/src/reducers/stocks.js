import * as types from '../actions/action-types';

const initalState = {
  data: {},
  loaded: false,
  filter: undefined
}

const stocks = (state = initalState, action = {}) => {
  switch (action.type) {
    case types.RECIVE_TRACKED_STOCKS:
      return {
        ...state,
        data: {
          ...state.data,
          ...action.payload.stocks
        },
        loaded: true
      };
    case types.SAVE_STOCK_INFO:
      return state;
    case types.UPDATE_FILTER:
      return {
        ...state,
        filter: action.payload.filterTerm
      }
    default:
      return state;
  }
}

export default stocks;
