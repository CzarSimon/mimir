/**
 * @file Handles all database interactions except connection, which is handled by ../server.js
 * @author Simon Lindgren, <simon.g.lindgren@gmail.com>
 */

'use strict';

const r = require('rethinkdb')
    , _ = require('lodash')
    , capitalize = require('lodash')['capitalize']
    , upperCase = require('lodash')['upperCase']
    , map = require('lodash')['map']
    , replace = require('lodash')['replace']
    , mapKeys = require('lodash')['mapKeys']
    , { dayType, nowUTC } = require('./helper-methods');

/* ---- API implementation ---- */

const fetch_stock_data = (tickers, conn, callback) => {
  r.table('stocks').filter(stock => {
    return r.expr(tickers).contains(stock('ticker'))
  }).without('id').run(conn, (err, res) => {
    const dt = dayType();
    const hour = nowUTC().getHours();
    res.toArray((err, res) => {
      callback(err, _.mapKeys(_pluck_stats(res, dt, hour), val => val.ticker));
    });
  });
}

const search_stocks = (query, conn, callback) => {
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

const insert_stock_data = (data, conn) => {
  const formated_data = _format_data(data);
  r.table('stocks').filter(stock => r.expr(formated_data.keys).contains(stock('ticker')))
  .run(conn, (err, res) => {
    if (err) throw err;
    res.each((err, res) => {
      if (err) throw err;
      r.table('stocks').get(res.id).update(
        Object.assign({}, res, formated_data.obj[res.ticker])
      ).run(conn, (err, res) => {
        if (err) throw err;
      })
    })
  });
}

const get_all_stocks = (conn, callback) => {
  r.table('stocks').without('id')
  .run(conn, (err, res) => {
    res.toArray((err, res) => {
      callback(err, res);
    })
  });
}

/* ---- Public API ---- */

module.exports = {
  search_stocks,
  fetch_stock_data,
  insert_stock_data,
  get_all_stocks
};

/* ---- Private functions ---- */

const _pluck_stats = (res, day_type, hour) => Object.assign({}, res, {
  stdev: res.stdev[day_type][hour],
  mean: res.mean[day_type][hour]
});

const _format_data = (data) => (
  {
    keys: _.map(data, (val, key) => _remove_dollar_tag(key)),
    obj: _.mapKeys(data, (val, key) => _remove_dollar_tag(key))
  }
)

const _remove_dollar_tag = (str) => _.replace(str, '$', '')
