/**
 * @file Handles all database interactions except connection, which is handled by ../server.js
 * @author Simon Lindgren, <simon.g.lindgren@gmail.com>
 */

'use strict';

const r = require('rethinkdb')
const { sendToClustering, pickClusterAttributes } = require('./clustering-client')

const insert_articles = (articles, conn) => {
  //sendToClustering(pickClusterAttributes(articles))
  r.table('articles').insert(articles).run(conn, (err, res) => {
    if (err) throw err;
  });
}

const update_article = (id, article_update, conn) => {
  //sendToClustering(pickClusterAttributes(article_update))
  r.table('articles').get(id).update(article_update).run(conn, (err, res) => {
    if (err) throw err;
  });
}

/* --- Returns the leader articles of the highest ranked clusters at a given date --- */
const fetchTopArticles = (ticker, top, date, conn, callback) => {
  r.table('articles').getAll(r.args(
    r.table('article_clusters').getAll(date, {index: 'date'})
    .filter({ticker: ticker})
    .filter(r.row('leader')('Score')('SubjectScore').gt(0))
    .orderBy(r.desc('score'))
    .limit(top)('leader')('UrlHash')
  ))
  .pluck('compound_score', 'url', 'timestamp', 'title', 'twitter_references', 'summary')
  .run(conn, (err, res) => {
    if (!err) {
      res.toArray((err, res) => {
        callback(err, res);
      })
    } else {
      console.log(err);
    }
  });
}

// Returns list of tickers if url present
const check_for_article = (url, conn, callback) => {
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
  check_for_article: check_for_article,
  insert_articles: insert_articles,
  update_article: update_article,
  fetchTopArticles: fetchTopArticles
};
