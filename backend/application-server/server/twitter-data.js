const r = require('rethinkdb');
const _ = require('lodash');
const { STOCK_TABLE } = require('./database');
const { nowUTC, dayType, isEmpty } = require('./helper-methods');

const HOURS_IN_DAY = 24;

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
* updateStockVolumes() Updates the mean and standard deviation values all stocks supplied
* Takes request and response objects and database connection as arguments
*/
const updateStockVolumes = (req, res, conn) => {
  console.log(req.body);
  const { volumes, hour } = req.body;
  if (!isEmpty(volumes) && !isEmpty(hour)) {
    insertStockData(hour, formatData(volumes), conn, err => {
      if (!err) {
        res.sendStatus(200);
      } else {
        console.log(err);
        res.sendStatus(500);
      }
    });
  } else {
    res.status(400).send('no data supplied');
  }
}

/**
* updateStockStats() Updates the mean and standard deviation values all stocks supplied
* Takes request and response objects and database connection as arguments
*/
const updateStockStats = (req, res, conn) => {
  console.log(req.body);
  if (_.size(req.body) > 0) {
    updateMeanAndStdev(req.body, conn, err => {
      if (err) {
        console.log(err);
        res.status(500).send(err);
      } else {
        res.sendStatus(200);
      }
    })
  } else {
    res.status(400).send('no data supplied');
  }
}

/**
* updateMeanAndStdev() Updates the mean and stdev data in the database for
* supplied stocks and respective data.
* Takes stockData, a database connection and a callback as arguments
*/
const updateMeanAndStdev = (stockData, conn, callback) => {
  const data = r.expr(stockData);
  r.table(STOCK_TABLE)
   .filter(stock => r.expr(_.keys(stockData)).contains(stock('ticker')))
   .update(stock => {
     return {
       mean: data.getField(stock('ticker')).getField('mean'),
       stdev: data.getField(stock('ticker')).getField('stdev')
     }
   }).run(conn, (err, res) => {
     console.log(res);
     callback(err);
   })
}

/**
* insertStockData() Inserts new twitter statistics about a list of stocks
* Takes a stock data object, a database connection and a callback as parameters
*/
const insertStockData = (hour, stockData, conn, callback) => {
  r.table(STOCK_TABLE).filter(stock => r.expr(_.keys(stockData)).contains(stock('ticker')))
  .run(conn, (err, res) => {
    if (err) {
      callback(err);
      return;
    }
    res.each((err, stock) => {
      if (err) {
        callback(err);
        return;
      }
      updateStockVolume(stock, hour, stockData[stock.ticker], conn)
    })
    callback(null);
    return;
  });
}

/**
* updateStockVolume() Inserts new twitter statistics about a list of stocks
* Takes a stock database object, an hour, a volume, a database connection as parameters.
* Updates the stock database object volume for the supplied hour.
*/
const updateStockVolume = (stock, hour, data, conn) => {
  if (hour < HOURS_IN_DAY && hour >= 0) {
    r.table(STOCK_TABLE)
    .get(stock.id)
    .update({
      volume: updateVolume(stock.volume, hour, data.volume),
      minute: data.minute
    })
    .run(conn, err => {
      if (err) {
        console.log(err);
      }
    })
  }
}

// updateVolume() Adds a new volume for the supplied hour to the array of volumes
const updateVolume = (volume, hour, newVolume) => {
  if (!validVolume(volume)) {
    volume = newVolumes();
  }
  volume[hour] = newVolume;
  return volume;
}

// validVolume() Chacks the supplied volume is of a correct format
const validVolume = volume => {
  const OBJECT_TYPE = 'object';
  return (!isEmpty(volume) && typeof(volume) === OBJECT_TYPE && volume.length === HOURS_IN_DAY);
}

// newVolumes() Creates an 24 item long volume array with 0 voluems for all entries
const newVolumes = () => _.fill(Array(HOURS_IN_DAY), 1);

// formatData() Downcases the keys of all nested objects
const formatData = data => _.mapValues(data, val => downCaseKeys(val))

// downCaseKeys() Downcases object keys
const downCaseKeys = object => _.mapKeys(object, (val, key) => _.toLower(key))

// Publicly exposed functions
module.exports = {
  getTwitterData,
  updateStockStats,
  updateStockVolumes
}
