'use strict';

import * as types from './action-types';
import { createAction } from 'redux-actions';

export const fetch_news_items = (ticker, socket) => {
  return (dispatch) => {
    socket.removeListener('DISPATCH_NEWS_ITEMS');
    socket.emit('FETCH_NEWS_ITEMS', { ticker });
    return socket.on('DISPATCH_NEWS_ITEMS', payload => {
      const news_items = { [ticker]: JSON.parse(payload.data) };
      dispatch(recive_news_items(news_items));
    });
  }
}

export const recive_news_items = createAction(types.RECIVE_NEWS_ITEMS, data => (
  { data }
))
