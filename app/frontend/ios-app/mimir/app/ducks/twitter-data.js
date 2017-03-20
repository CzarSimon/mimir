'use strict'
import { createAction } from 'redux-actions';
import { getTwitterData } from '../methods/server/twitter-miner';
import { mapKeys, replace } from 'lodash';

/* --- Types --- */
export const RECIVE_TWITTER_DATA = 'RECIVE_TWITTER_DATA'
export const FETCH_TWITTER_DATA = 'FETCH_TWITTER_DATA' //Add failure scenario

const initialState = {
  data: [],
  loaded: false
}

/* --- Reducer --- */
export default twitterData = (state = initialState, action = {}) => {
  switch (action.type) {
    case RECIVE_TWITTER_DATA:
      return {
        data: mapKeys(action.payload.data, (val, key) => replace(key, '$', '')),
        loaded: true
      }
    default:
      return state;
  }
}

/* --- Actions --- */
export const fetchTwitterData = (user, socket) => (() => socket.emit(FETCH_TWITTER_DATA, { user }))

export const reciveTwitterData = createAction(RECIVE_TWITTER_DATA, data => {({ data }))
