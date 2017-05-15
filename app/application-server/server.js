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
  database.get_all_stocks(app._rdb_conn, (error, result) => {
    if (error) {
      console.log(error.message);
    } else {
      res.send({ list: result, date: nowUTC()})
    }
  });
});

app.post('/stockList', (req, res) => {
  const dict = req.body;
  if (dict) {
    sockets.alert_new_twitter_data(io.sockets, dict, app._rdb_conn);
    sockets.update_stocklist(io.sockets, app._rdb_conn);
    res.send('success');
  } else {
    res.send('failure');
  }
});

app.post('/tweet_volumes', (req, res) => {
  const dict = req.body;
  if (dict) {
    sockets.alert_new_twitter_data(io.sockets, dict, app._rdb_conn);
    sockets.update_stocklist(io.sockets, app._rdb_conn);
    res.send('success');
  } else {
    res.send('failure');
  }
});

//No need for real time updateing socket does not nees to be involved
app.post('/mean_and_stdev', (req, res) => routes.update_stock_stats(req, res, app._rdb_conn));

//app.get('/get_stock_data', (req,res) => routes.get_stock_data(req, res, app._rdb_conn));


const start_express = (connection) => {
  app._rdb_conn = connection;
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
  start_express(conn);
});
