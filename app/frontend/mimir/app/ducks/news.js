'use strict'
import { createAction } from 'redux-actions';
import { getRequest } from '../methods/api-methods';

/* --- Types --- */
export const FETCH_NEWS_ITEMS = 'mimir/news/FETCH';
export const RECIVE_NEWS_FAILURE = 'mimir/news/RECIVE_FAIL';
export const RECIVE_NEWS_ITEMS = 'mimir/news/RECIVE';

/* --- Reducer --- */
const news = (state = {}, action = {}) => {
  switch (action.type) {
    case RECIVE_NEWS_ITEMS:
      return {
        ...state,
        ...action.payload.data
      };
    default:
      return state;
  }
}
export default news;

/* --- Actions --- */
export const fetchNewsItems = (ticker, period) => (
  dispatch => (
    getRequest(`api/news/${ticker}/5/${period}`)
    .then(news => dispatch(reciveNewsItems({ [ticker]: news })))
    .catch(err => {
      console.log(err);
    })
  )
);

export const reciveNewsItems = createAction(RECIVE_NEWS_ITEMS, data => ({ data }));
