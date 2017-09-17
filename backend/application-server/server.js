/**
 * @file Application server for mimir iOS-app and temporary frontend.
 * @author Simon Lindgren, <simon.g.lindgren@gmail.com>
 */

'use strict'

const config = require('./config');
const routes = require('./server/routes');
const database = require('./server/database');
const { nowUTC } = require('./server/helper-methods');

const express = require('express');
const path = require('path');
const bodyParser = require('body-parser');

const app = express();
const server = require('http').createServer(app);
const r = require('rethinkdb');

app.use(express.static(path.join(__dirname, 'public')));
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

const startExpress = connection => {
  app.rdb = connection;
  app.use('/', routes.setupRoutes(app.rdb));
  server.listen(config.express.port, () => {
    console.log("Server running on port: " + config.express.port);
  });
}

r.connect(config.rethinkdb, (err, conn) => {
  if (err) {
    console.log("Could not establish connection with database");
    throw err.message;
  }
  startExpress(conn);
});
