/**
 * @file Handles socket interactions after connection.
 * @author Simon Lindgren, <simon.g.lindgren@gmail.com>
 */

'use strict';

const events = require('./events');
const database = require('./database');
const { nowUTC } = require('./helper-methods');
const config = require('../config');
const request = require('request');
const _ = require('lodash');


const newsData = socket => {
  socket.on(events.FETCH_NEWS_ITEMS, payload => {
    const ticker = _.upperCase(payload.ticker);
    const { address, port } = config.news_server;
    const fetchAddress = `${address}:${port}/api/news/${ticker}/5`;
    request(fetchAddress, (error, response, body) => {
      if (!error && response.statusCode === 200) {
        socket.emit(events.DISPATCH_NEWS_ITEMS, { data: body });
      } else {
        console.log(response.statusCode);
        console.log(error.message);
      }
    });
  });
}

const clientInfo = socket => {
  socket.emit(events.GET_CLIENT_INFO, "GET INFO");
  socket.on(events.DISPATCH_CLIENT_INFO, data => {
    console.log("New user made connection on: " + data.client_machine);
  });
}

const stockData = (socket, conn) => {
  socket.on(events.FETCH_TWITTER_DATA, payload => {
    const tickers = payload.user.tickers;
    if (tickers.length) {
      database.fetchStockData(tickers, conn, (err, res) => {
        if (err) {
          console.log(events.TWITTER_DATA_FAILURE);
          socket.emit(events.TWITTER_DATA_FAILURE, { data: null, error: err.message });
        } else {
          socket.emit(events.DISPATCH_TWITTER_DATA, { data: res, error: null });
        }
      })
    }
  });
}

const searchStocks = (socket, conn) => {
  socket.on(events.FETCH_SEARCH_RESULTS, payload => {
    database.searchStocks(payload.query, conn, (err, res) => {
      if (err) {
        socket.emit(events.DISPATCH_SEARCH_FAILURE, { results: null, error: err.message });
      } else {
        socket.emit(events.DISPATCH_SEARCH_RESULTS, { results: res, error: null });
      }
    })
  });
}

const alertNewTwitterData = (sockets, data, conn) => {
  database.insertStockData(data, conn);
  sockets.emit(events.NEW_TWITTER_DATA);
}

const updateStocklist = (sockets, conn) => {
  database.getAllStocks(conn, (err, res) => {
    if (err) {
      console.log("In update stocklist:", err.message);
    } else {
      sockets.emit(events.UPDATE_STOCKLIST, { list: res, date: nowUTC })
    }
  });
}

module.exports = {
  newsData,
  clientInfo,
  stockData,
  searchStocks,
  alertNewTwitterData,
  updateStocklist
}
