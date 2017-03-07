/**
 * @file News server for mimir, ranks articles and fetches stored ranked ones upon request.
 * @author Simon Lindgren, <simon.g.lindgren@gmail.com>
 */

'use strict';
//require('./server/memory-analysis');

const express = require('express')
    , bodyParser = require('body-parser')
    , config = require('./config')
    , r = require('rethinkdb')
    , { get_news_articles, rank_articles } = require('./server/routes')
    , app = express()
    , server = require('http').createServer(app);

app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

app.get('/news/:ticker/:top', (req, res) => { get_news_articles(req, res, app._rdb_conn) });
app.post('/rankArticle', (req, res) => { rank_articles(req, res, app._rdb_conn) });


const start_express = (connection) => {
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
  start_express(conn);
})
