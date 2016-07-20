import { map } from 'lodash';
import * as types from './../actions/action-types';

const initial_state = {
  data: [],
  loaded: false
};

const stocks = (state = initial_state, action = {}) => {
  switch (action.type) {
    case types.RECIVE_STOCK_DATA:
      console.log(action.payload.data);
      return {
        ...state,
        data: {
          ...state.data,
          ...action.payload.data
        },
        loaded: true
      };
    default:
      return state;
  }
}

export default stocks;

/*
case types.UPDATE_STOCK_DATA:
  return {
    ...state,
    data: map(state.data, (val, key) => ({
      ...val,
      ...action.payload.data[key]
    })),
    loaded: true
  };
*/
