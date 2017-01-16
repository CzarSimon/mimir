import * as types from './action-types';
import { createAction } from 'redux-actions';

export const reciveUntrackedTickers =
  createAction(types.RECIVE_UNTRACKED_TICKERS, tickers => (
    { tickers }
  ))

export const fetchUntrackedTickers = () => {
  return dispatch => {
    return fetch("http://localhost:8000/untracked-tickers")
    .then(res => res.json())
    .then(tickers => dispatch(reciveUntrackedTickers(tickers)))
    .catch(err => console.log("Error in fetch tickers: ", err))
  }
}
