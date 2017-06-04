'use strict';

const pg = require('pg');
const config = require('../config');
const { map } = require('lodash');

const setupConnection = () => {
  const pool = new pg.Pool(config.postgres);
  //client.connect();
  return client;
}

const getSearchSugestions = (tickers, client, callback) => {
  const sql = "SELECT ticker, name FROM stocks WHERE is_tracked=TRUE AND ticker NOT IN ("
              + constructParams(tickers) + ") ORDER BY total_count DESC LIMIT 5";
  client.query(sql, tickers, (err, result) => callback(err, result));
}

const constructParams = tickers => map(tickers, (_, index) => '$' + (index + 1))

module.exports = {
  setupConnection,
  getSearchSugestions
}
