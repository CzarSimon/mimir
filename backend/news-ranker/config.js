'use strict';

const path = require('path');

module.exports = {
  rethinkdb: {
    host: process.env.db_host || 'localhost',
    port: 28015,
    db: 'mimir_news_db'
  },
  express: {
    port: 5000
  },
  rank_script: {
    path: path.join(__dirname, 'article_ranker'),
    name: "scrape_and_rank.pyc",
    command: "python3",
    twitter_users: 328000000.0,
    reference_weight: 800.0
  },
  clusterServer: {
    host: process.env.cluster_host || 'localhost',
    port: process.env.cluster_port || '6000',
    protocol: 'http'
  }
};
