/**
 * @file Handles socket interactions after connection.
 * @author Simon Lindgren, <simon.g.lindgren@gmail.com>
 */

'use strict';

const events = require('./events')
    , database = require('./database')
    , helper_methods = require('./helperMethods');

module.exports = {

  client_info: (socket) => {
    socket.emit(events.GET_CLIENT_INFO, "GET INFO");
    socket.on(events.DISPATCH_CLIENT_INFO, data => {
      console.log("New user made connection on: " + data.client_machine);
    });
  },

  stock_data: (socket, conn) => {
    socket.on(events.FETCH_TWITTER_DATA, payload => {
      const user_tickers = payload.user.tickers;
      if (user_tickers.length) {
        database.fetch_stock_data(user_tickers, conn, (err, res) => {
          if (err) {
            socket.emit(events.TWITTER_DATA_FAILURE, { data: null, error: err.message });
          } else {
            socket.emit(events.DISPATCH_TWITTER_DATA, { data: res, error: null });
          }
        })
      }
    });
  },

  search_stocks: (socket, conn) => {
    socket.on(events.FETCH_SEARCH_RESULTS, payload => {
      database.search_stocks(payload.query, conn, (err, res) => {
        if (err) {
          socket.emit(events.DISPATCH_SEARCH_FAILURE, { results: null, error: err.message });
        } else {
          socket.emit(events.DISPATCH_SEARCH_RESULTS, { results: res, error: null });
        }
      })
    });
  },

  alert_new_twitter_data: (sockets, data, conn) => {
    database.insert_stock_data(data, conn);
    sockets.emit(events.NEW_TWITTER_DATA);
  },

  update_stocklist: (sockets, conn) => {
    database.get_all_stocks(conn, (err, res) => {
      if (err) {
        console.log("In update stocklist:", err.message);
      } else {
        sockets.emit(events.UPDATE_STOCKLIST, { list: res, date: helper_methods.getDate() })
      }
    });
  }
}
