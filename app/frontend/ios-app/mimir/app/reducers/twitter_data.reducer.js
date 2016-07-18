import { mapKeys, replace } from 'lodash';
import * as types from './../actions/action-types';

const initial_state = {
  data: [],
  loaded: false
};

const twitter_data = (state = initial_state, action = {}) => {
  switch (action.type) {
    case types.RECIVE_TWITTER_DATA:
      return {
        data: mapKeys(action.payload.data, (val, key) => replace(key, '$', '')),
        loaded: true
      }
    default:
      return state;
  }
}

export default twitter_data;
