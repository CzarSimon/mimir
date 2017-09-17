const r = require("rethinkdb");
const { STOCK_TABLE } = require('./database');
const { isEmpty } = require('./helper-methods');
const _ = require('lodash');

/**
* getDescription() Retrives the stored description of a company
* Takes request and response objects and database connection as paramters
*/
const getDescription = (req, res, conn) => {
  const ticker = req.query.ticker;
  if (isEmpty(ticker)) {
    res.status(400).send("No user id supplied");
    return;
  }
  queryForDescription(ticker, conn, (err, result) => {
    if (!err) {
      if (isEmpty(result.description)){
        res.status(404).send("description not found")
      } else {
        res.status(200).send(result);
      }
    } else {
      console.log(err);
      res.status(500).send("could not retrive description");
    }
  })
}

/**
* queryForDescription() Retrives the description of a company form the database
* Takes a ticker string, database connection and a callback as parameters
*/
const queryForDescription = (ticker, conn, callback) => {
  r.table(STOCK_TABLE)
    .filter({ticker: ticker})
    .pluck('description')
    .run(conn, (err, res) => {
      if (err) {
        callback(err, {})
      }
      res.toArray((err, description) => {
        if (err) {
          callback(err, {})
        } else {
          callback(err, parseDescription(description));
        }
      })
    })
}

const parseDescription = description => {
  const firstDesc = _.head(description);
  return (!isEmpty(firstDesc)) ? firstDesc : {};
}

module.exports = {
  getDescription
};
