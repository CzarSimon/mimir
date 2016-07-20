'use strict';
import * as types from './action-types';
import { createAction } from 'redux-actions';
import { retrive_object} from './../methods/async-storage';
import { recive_user } from './user.actions';
import { fetch_twitter_data } from './twitter_data.actions';
import { fetch_stock_data } from './stock.actions';

export const logon_user = (socket) => {
  return (dispatch) => {
    return retrive_object("user")
    .then(user => {
      dispatch(recive_user(user))
      dispatch(fetch_twitter_data(user, socket))
      dispatch(fetch_stock_data(user.tickers))
    })
  }
}
