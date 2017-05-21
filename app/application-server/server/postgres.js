'use strict';

const pg = require('pg');
const config = require('../config');

const setupPostgresConn = () => {
  return new pg.Pool(config.postgres);
}

const getSearchSugestions = (tickers, pool, callback) => {
  const query = "SELECT ticker FROM TWEET_COUNT WHERE ticker NOT IN ($1::text) order by TWEET_COUNT DESC LIMIT 10";
  pool.query(query, tickers, (err, result) => {
    if (err) {
      return console.log(err);
    }
    callback(result);
  });
}
