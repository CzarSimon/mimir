'use strict';
import * as types from './action-types';
import { createAction } from 'redux-actions';
import socket from '../methods/server/socket';

export const toggle_search_active = createAction(types.TOGGLE_SEARCH_ACTIVE);

export const recive_search_results = createAction(types.RECIVE_SEARCH_RESULTS, results => (
  { results }
))

export const fetch_search_results = query => (() =>
  socket.emit(types.FETCH_SEARCH_RESULTS, { query })
)
