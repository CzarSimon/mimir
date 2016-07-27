'use strict';
import * as types from './action-types';
import { createAction } from 'redux-actions';
import { retrive_object, persist_object } from '../methods/async-storage';
import { generate_new_user } from '../methods/user';
import { recive_user, create_new_user } from './user.actions';
import { fetch_twitter_data } from './twitter_data.actions';
import { fetch_stock_data } from './stock.actions';

export const logon_user = (socket) => {
  return (dispatch) => {
    return retrive_object("user")
    .then(user => {
      const tickers = user.tickers;
      dispatch(recive_user(user))
      dispatch(fetch_twitter_data(user, socket))
      dispatch(fetch_stock_data(tickers))
    })
    .catch(err => {
      const new_user = generate_new_user();
      persist_object("user", new_user);
      dispatch(create_new_user(new_user))
      dispatch(fetch_twitter_data(new_user, socket))
      dispatch(fetch_stock_data(new_user.tickers))
    })
  }
}
