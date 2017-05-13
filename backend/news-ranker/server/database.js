/**
 * @file Handles all database interactions except connection, which is handled by ../server.js
 * @author Simon Lindgren, <simon.g.lindgren@gmail.com>
 */

'use strict';

const r = require('rethinkdb')
const { sendToClustering, pickClusterAttributes } = require('./clustering-client')

const insertArticles = (articles, conn) => {
  sendToClustering(pickClusterAttributes(articles))
  r.table('articles').insert(articles).run(conn, (err, res) => {
    if (err) throw err;
  });
}

const updateArticle = (id, article_update, conn) => {
  sendToClustering(pickClusterAttributes(article_update))
  r.table('articles').get(id).update(article_update).run(conn, (err, res) => {
    if (err) throw err;
  });
}


// Returns list of tickers if url present
const checkForArticle = (url, conn, callback) => {
  r.table('articles').getAll(url, {index: 'url'})
  .run(conn, (err, res) => {
    if (err) {
      console.log(conn);
      throw err;
    }
    res.toArray((err, res) => {
      callback(err, res);
    })
  })
}

module.exports = {
  insert_articles: insertArticles,
  update_article: updateArticle,
  fetchTopArticles: fetchTopArticles
};
