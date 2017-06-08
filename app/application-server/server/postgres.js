'use strict';

const pg = require('pg');
const config = require('../config');
const { map } = require('lodash');

const setupConnection = () => {
  const pool = new pg.Pool(config.postgres);
  return pool;
}

const getSearchSugestions = (tickers, pool, callback) => {
  const sql = getSQL(tickers);
  if (tickers.length) {
    pool.query(sql, tickers, (error, result) => callback(error, result));
  } else {
    pool.query(sql, (error, result) => callback(error, result));
  }
}

const getSQL = tickers => (
  (tickers.length)
  ? "SELECT ticker, name FROM stocks WHERE is_tracked=TRUE AND ticker NOT IN (" + constructParams(tickers) + ") ORDER BY total_count DESC LIMIT 5"
  : "SELECT ticker, name FROM stocks WHERE is_tracked=TRUE ORDER BY total_count DESC LIMIT 5"
)

const constructParams = tickers => map(tickers, (_, index) => '$' + (index + 1))

module.exports = {
  setupConnection,
  getSearchSugestions
}
