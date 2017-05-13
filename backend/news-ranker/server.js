/**
 * @file News server for mimir, ranks articles and fetches stored ranked ones upon request.
 * @author Simon Lindgren, <simon.g.lindgren@gmail.com>
 */

'use strict';

const express = require('express');
const bodyParser = require('body-parser');
const config = require('./config');
const r = require('rethinkdb');
const { rankArticles, handleRankedArticle } = require('./server/routes');
const app = express();
const server = require('http').createServer(app);

app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

app.post('/rankArticle', (req, res) => { rankArticles(req, res, app._rdb_conn) });
app.post('/ranked-article', (req, res) => { handleRankedArticle(req, res, app._rdb_conn) });


const startExpress = (connection) => {
  app._rdb_conn = connection
  server.listen(config.express.port, () => {
    console.log("Server running on http://localhost:" + config.express.port);
  })
}

r.connect(config.rethinkdb, (err, conn) => {
  console.log("Attempting start at: " + new Date().toUTCString());
  if (err) {
    throw err.message
  }
  startExpress(conn);
})
