'use strict';
import { map } from 'lodash';
import * as types from './../actions/action-types';

const initial_state = {};

const descriptions = (state = initial_state, action = {}) => {
  switch (action.type) {
    case types.RECIVE_COMPANY_DESC:
      return {
        ...state,
        [action.payload.ticker]: action.payload.description
      };
    case types.RECIVE_DESC_FAILURE:
      return {
        ...state,
        [action.payload.ticker]: action.payload.error
      };
    default:
      return state;
  }
}

export default descriptions;
