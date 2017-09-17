/**
 * @file Handles all database interactions except connection, which is handled by ../server.js
 * @author Simon Lindgren, <simon.g.lindgren@gmail.com>
 */

'use strict';

const r = require('rethinkdb');
const _ = require('lodash');
const { dayType, nowUTC } = require('./helper-methods');

// Table names
const USER_TABLE = "users";
const STOCK_TABLE = "stocks";


/* ---- API implementation ---- */

const fetchStockData = (tickers, conn, callback) => {
  r.table('stocks').filter(stock => {
    return r.expr(tickers).contains(stock('ticker'))
  }).without('id').run(conn, (err, res) => {
    const hour = nowUTC().getHours();
    res.toArray((err, res) => {
      const trimedData = _.map(res, stockData => pluckStats(stockData, dayType(), hour))
      callback(err, _.mapKeys(trimedData, val => val.ticker));
    });
  });
}

const getAllStocks = (conn, callback) => {
  r.table('stocks').without('id')
  .run(conn, (err, res) => {
    res.toArray((err, res) => {
      callback(err, res);
    })
  });
}

/* ---- Public API ---- */

module.exports = {
  fetchStockData,
  getAllStocks,
  USER_TABLE,
  STOCK_TABLE
};

/* ---- Private functions ---- */

const pluckStats = (data, dayType, hour) => {
  return Object.assign({}, data, {
    stdev: data.stdev[dayType][hour],
    mean: data.mean[dayType][hour]
  });
}

const removeDollarTag = (str) => _.replace(str, '$', '')
