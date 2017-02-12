import * as types from '../actions/action-types';
import _ from 'lodash';

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
    case types.REMOVE_UNTACKED_STOCK:
        return {
          ...state,
          data: _.omit(state.data, [action.payload.ticker])
        }
    default:
      return state;
  }
}

export default stocks;
