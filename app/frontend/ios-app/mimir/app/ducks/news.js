'use strict'
import { createAction } from 'redux-actions';

/* --- Types --- */
export const FETCH_NEWS_ITEMS = 'FETCH_NEWS_ITEMS'
export const RECIVE_NEWS_FAILURE = 'RECIVE_NEWS_FAILURE'
export const RECIVE_NEWS_ITEMS = 'RECIVE_NEWS_ITEMS'
export const DISPATCH_NEWS_ITEMS = 'DISPATCH_NEWS_ITEMS'

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
export default news

/* --- Actions --- */
export const fetchNewsItems = (ticker, socket) => (
  dispatch => {
    socket.removeListener(DISPATCH_NEWS_ITEMS);
    socket.emit(FETCH_NEWS_ITEMS, { ticker });
    return socket.on(DISPATCH_NEWS_ITEMS, payload => {
      const newsItems = { [ticker]: JSON.parse(payload.data) };
      dispatch(reciveNewsItems(newsItems));
    });
  }
)

export const reciveNewsItems = createAction(RECIVE_NEWS_ITEMS, data => ({ data }))
