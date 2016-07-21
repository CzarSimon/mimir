import { map, mapValues } from 'lodash';
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
    case types.RECIVE_UPDATED_STOCK_DATA:
      return {
        ...state,
        data: mapValues(state.data, (val) => (
          {
            ...val,
            ...action.payload.data[val.Symbol]
          }
        )),
        loaded: true
      };
    default:
      return state;
  }
}

export default stocks;
