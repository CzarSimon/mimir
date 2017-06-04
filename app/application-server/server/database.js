/**
 * @file Handles all database interactions except connection, which is handled by ../server.js
 * @author Simon Lindgren, <simon.g.lindgren@gmail.com>
 */

'use strict';

const r = require('rethinkdb');
const _ = require('lodash');
const { dayType, nowUTC } = require('./helper-methods');

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

const searchStocks = (query, conn, callback) => {
  r.table('stocks').filter(stock => {
    return (
      stock("name").match("^" + _.capitalize(query))
      .or(stock("ticker").match("^" + _.upperCase(query)))
    );
  }).without('id').run(conn, (err, res) => {
    if (err) throw err;
    res.toArray((err, res) => {
      callback(err, res);
    })
  });
}

const insertStockData = (data, conn) => {
  const formatedData = formatData(data);
  r.table('stocks').filter(stock => r.expr(_.keys(formatedData)).contains(stock('ticker')))
  .run(conn, (err, res) => {
    if (err) throw err;
    res.each((err, res) => {
      if (err) throw err;
      r.table('stocks').get(res.id).update(
        Object.assign({}, res, formatedData[res.ticker])
      ).run(conn, (err, res) => {
        if (err) throw err;
      })
    })
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
  searchStocks,
  fetchStockData,
  insertStockData,
  getAllStocks
};

/* ---- Private functions ---- */

const pluckStats = (data, dayType, hour) => {
  return Object.assign({}, data, {
    stdev: data.stdev[dayType][hour],
    mean: data.mean[dayType][hour]
  });
}

const formatData = data => _.mapValues(data, val => downCaseKeys(val))

const downCaseKeys = object => _.mapKeys(object, (val, key) => _.toLower(key))

const removeDollarTag = (str) => _.replace(str, '$', '')
