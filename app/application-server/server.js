/**
 * @file Application server for mimir iOS-app and temporary frontend.
 * @author Simon Lindgren, <simon.g.lindgren@gmail.com>
 */

'use strict'

const config = require('./config');
const sockets = require('./server/sockets');
const events = require('./server/events');
const routes = require('./server/routes');
const database = require('./server/database');
const postgres = require('./server/postgres');
const { nowUTC } = require('./server/helper-methods');

const express = require('express');
const path = require('path');
const bodyParser = require('body-parser');

const app = express();
const server = require('http').createServer(app);
const io = require('socket.io').listen(server);
const r = require('rethinkdb');

app.use(express.static(path.join(__dirname, 'public')));
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

app.get('/stockList', (req, res) => {
  database.getAllStocks(app.rdb, (error, result) => {
    if (error) {
      console.log(error.message);
    } else {
      res.send({ list: result, date: nowUTC()})
    }
  });
});

app.post('/tweet_volumes', (req, res) => {
  const dict = req.body;
  if (dict) {
    sockets.alertNewTwitterData(io.sockets, dict, app.rdb);
    sockets.updateStocklist(io.sockets, app.rdb);
    res.sendStatus(200);
  } else {
    res.sendStatus(500);
  }
});

app.post('/mean_and_stdev', (req, res) =>
  routes.updateStockStats(req, res, app.rdb));

app.get('/search-sugestions', (req, res) =>
  routes.getSearchSugestions(req, res, app.pg));


const startExpress = connection => {
  app.rdb = connection;
  app.pg = postgres.setupConnection();
  server.listen(config.express.port, () => {
    console.log("Server running on port: " + config.express.port);
  });
}

r.connect(config.rethinkdb, (err, conn) => {
  if (err) {
    console.log("Could not establish connection with database");
    throw err.message;
  }
  io.on(events.CONNECTION, socket => {
    sockets.client_info(socket);
    sockets.stock_data(socket, conn);
    sockets.search_stocks(socket, conn);
    sockets.news_data(socket);
  });
  startExpress(conn);
});
