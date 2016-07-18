import { map } from 'lodash';
import * as types from './../actions/action-types';

const initial_state = {
  data: [],
  loaded: false
};

const stocks = (state = initial_state, action = {}) => {
  switch (action.type) {
    case types.RECIVE_STOCK_DATA:
      return {
        ...state,
        data: action.payload.data,
        loaded: true
      };
    default:
      return state;
  }
}

export default stocks;
