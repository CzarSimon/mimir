const r = require('rethinkdb');
const _ = require('lodash');
const { STOCK_TABLE } = require('./database');
const { nowUTC, dayType, isEmpty } = require('./helper-methods');

/**
* getTwitterData() Gets volume, mean and stdev data for the supplied tickers
* Takes request and response objects and database connection as paramters
*/
const getTwitterData = (req, res, conn) => {
  const tickers = parseTickers(req.query)
  if (!tickers.length) {
    res.status(400).send("no tickers supplied");
    return;
  }
  queryForTickers(tickers, conn, (err, twitterData = {}) => {
    if (!err) {
      res.status(200).send(twitterData);
    } else {
      console.log(err);
      res.sendStatus(500);
    }
  });
}

/**
* queryForTickers() Retrives a specified list of twitter data from the database
* Takes a ticker array, database connection and a callback as parameters
*/
const queryForTickers = (tickers, conn, callback) => {
  const stockFields = ['ticker', 'minute', 'volume', 'mean', 'stdev']
  r.table(STOCK_TABLE).filter(stock => {
    return r.expr(tickers).contains(stock('ticker'))
  })
  .pluck(stockFields)
  .run(conn, (err, res) => {
    if (err) {
      callback(err);
      return;
    }
    res.toArray((err, res) => {
      if (err) {
        callback(err);
        return;
      }
      callback(err, parseTwitterData(res));
    })
  })
}

// parseTwitterData() Alters the mean & stdev field to only have one daytype
const parseTwitterData = twitterData => {
  const dt = dayType();
  const hour = nowUTC().getHours();
  return _.map(twitterData, stock => (
    Object.assign({}, stock, {
      mean: stock.mean[dt],
      stdev: stock.stdev[dt],
      hour
    })
  ));
}

// parseTickers() Parses the tickers in a request an returns as an array
const parseTickers = query => {
  tickers = query.ticker
  switch (typeof(tickers)) {
    case "object":
      return tickers;
    case "string":
      return [ tickers ];
    default:
      return []
  }
}

/**
* updateStockStats() Updates the mean and standard deviation values all stocks supplied
* Takes request and response objects and database connection as arguments
*/
const updateStockStats = (req, res, conn) => {
  const stockStats = req.body;
  console.log(stockStats);
  if (!isEmpty(stockStats)) {
    insertStockData(formatData(stockStats), conn, err => {
      if (!err) {
        console.log("no error");
        res.sendStatus(200);
      } else {
        console.log("an error");
        console.log(err);
        res.sendStatus(500);
      }
    });
  } else {
    res.status(400).send('no data supplied');
  }
}

/**
* insertStockData() Inserts new twitter statistics about a list of stocks
* Takes a stock data object, a database connection and a callback as parameters
*/
const insertStockData = (stockData, conn, callback) => {
  r.table(STOCK_TABLE).filter(stock => r.expr(_.keys(stockData)).contains(stock('ticker')))
  .run(conn, (err, res) => {
    if (err) {
      callback(err);
      return;
    }
    res.each((err, res) => {
      if (err) {
        callback(err);
        return;
      }
      r.table(STOCK_TABLE).get(res.id).update(
        Object.assign({}, res, stockData[res.ticker])
      ).run(conn, (err, res) => {
        if (err) {
          callback(err);
          return;
        }
      })
    })
    callback(null);
    return
  });
}

// formatData() Downcases the keys of all nested objects
const formatData = data => _.mapValues(data, val => downCaseKeys(val))

// downCaseKeys() Downcases object keys
const downCaseKeys = object => _.mapKeys(object, (val, key) => _.toLower(key))

// Publicly exposed functions
module.exports = {
  getTwitterData,
  updateStockStats
}
