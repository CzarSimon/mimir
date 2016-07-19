'use strict';
import * as types from './action-types';
import { createAction } from 'redux-actions';

export const set_active_ticker = createAction(types.SET_ACTIVE_TICKER, (ticker) => (
  { ticker }
))

export const select_tab = createAction(types.SELECT_TAB, (tab) => (
  { tab }
))
