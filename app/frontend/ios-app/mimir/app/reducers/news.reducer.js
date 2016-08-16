'use strict';

import * as types from '../actions/action-types';

const initial_state = {};

const news = (state = initial_state, action = {}) => {
  switch (action.type) {
    case types.RECIVE_NEWS_ITEMS:
      return {
        ...state,
        ...action.payload.data
      };
    default:
      return state;
  }
}

export default news;
