const routes = require('express').Router();
const database = require('./database');
const twitterData = require('./twitter-data');
const user = require('./user');

/**
* setupRoutes() Sets up routes
* Takes a database connection as argument
*/
const setupRoutes = conn => {
  twitterDataRoutes(conn);
  userRoutes(conn);
  return routes;
}

/**
* twitterDataRoutes() sets up routes for dealing with twitter data
* Takes a database connection as argument
*/
const twitterDataRoutes = conn => {
  routes.get('/api/app/twitter-data',
    (req, res) => twitterData.getTwitterData(req, res, conn));
  routes.post('/api/app/twitter-data/volumes',
    (req, res) => twitterData.updateStockStats(req, res, conn));
  routes.post('/api/app/twitter-data/mean-and-stdev',
    (req, res) => twitterData.updateStockStats(req, res, conn));
}

/**
* userRoutes() sets up routes for dealing with users
* Takes a database connection as argument
*/
const userRoutes = conn => {
  routes.post('/api/app/user',
    (req, res) => user.newUser(req, res, conn));
  routes.get('/api/app/user',
    (req, res) => user.getUser(req, res, conn));
  routes.post('/api/app/user/session',
    (req, res) => user.recordSession(req, res, conn));
  routes.post('/api/app/user/search',
    (req, res) => user.saveSearch(req, res, conn));
  routes.post('/api/app/user/ticker',
    (req, res) => user.addTicker(req, res, conn));
  routes.delete('/api/app/user/ticker',
    (req, res) => user.deleteTicker(req, res, conn));
}

/**
* getStockData() Fetches stock data for all supplied tickers
* Takes request and response objects and database connection as arguments
*/
const getStockData = (req, res, conn) => {
  tickers = req.body.tickers;
  if (tickers.length) {
    database.fetchStockData(tickers, conn, (err, res) => {
      if (err) {
        res.status(500).send({data: null, error: err.message});
      } else {
        res.status(200).send({data: res, error: null})
      }
    })
  }
}

module.exports = {
  setupRoutes,
  getStockData
}
