import * as types from './action-types';
import { createAction } from 'redux-actions';
import { browserHistory } from 'react-router';
import { createHttpObject } from '../methods/helper-methods';


export const reciveUserCredentials =
  createAction(types.RECIVE_USER_CREDENTIALS, (username, token) => (
    {username, token}
  ))


export const loginUser = (user, pwd) => {
  const httpObject = createHttpObject('POST', '', {user, pwd})
  return dispatch => {
    fetch('http://localhost:8000/login', httpObject)
    .then(res => res.json())
    .then(res => {
      if (res.Token) {
        dispatch(reciveUserCredentials(res.Username, res.Token))
        browserHistory.push('/')
      } else {
        console.log(res.Response)
      }
    })
  }
}
