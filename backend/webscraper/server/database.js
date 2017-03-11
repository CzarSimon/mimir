/**
 * @file Handles all database interactions except connection, which is handled by ../server.js
 * @author Simon Lindgren, <simon.g.lindgren@gmail.com>
 */

'use strict';

const r = require('rethinkdb')
const { sendToClustering, pickClusterAttributes } = require('./clustering-client')

const insert_articles = (articles, conn) => {
  sendToClustering(pickClusterAttributes(articles))
  r.table('articles').insert(articles).run(conn, (err, res) => {
    if (err) throw err;
  });
}

const update_article = (id, article_update, conn) => {
  sendToClustering(pickClusterAttributes(article_update))
  r.table('articles').get(id).update(article_update).run(conn, (err, res) => {
    if (err) throw err;
  });
}

const fetch_top_articles = (ticker, top, date, conn, callback) => {
  r.table('articles').getAll(date, {index: 'timestamp'})
  .filter(article => article("compound_score").hasFields(ticker))
  .filter(article => article("compound_score")(ticker).gt(article("reference_score")))
  .orderBy(r.desc(article => article("compound_score")(ticker)))
  .limit(top)
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
  fetch_top_articles: fetch_top_articles
};
