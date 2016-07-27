import * as types from './action-types';
import { createAction } from 'redux-actions';
import { retrive_object} from './../methods/async-storage';

export const fetch_user = () => {
  return (dispatch) => {
    return retrive_object("user")
    .then(user => dispatch(recive_user(user)))
  }
}

export const recive_user = createAction(types.RECIVE_USER, user => (
  { user }
))

export const create_new_user = createAction(types.CREATE_NEW_USER, user => (
  { user }
))

export const add_ticker = createAction(types.ADD_TICKER, (ticker) => (
  { ticker }
))

export const remove_ticker = createAction(types.REMOVE_TICKER, (ticker) => (
  { ticker }
))

export const toggle_modifiable = createAction(types.TOGGLE_MODIFIABLE)
