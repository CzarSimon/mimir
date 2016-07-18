import * as types from './action-types';
import { createAction } from 'redux-actions';
import { get_twitter_data } from '../methods/server/twitter-miner';

export const fetch_twitter_data = (user, socket) => (() => 
  socket.emit('GET TWITTER DATA', { user })
);

export const recive_twitter_data = createAction(types.RECIVE_TWITTER_DATA, data => (
  { data }
))
