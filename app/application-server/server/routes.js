'use strict';

const request = require('request')
    , config = require('../config')
    , _ = require('lodash');

const get_news = (req, res) => {
  const ticker = _.upperCase(req.params.ticker)
      , { address, port } = config.news_server
      , fetch_address = `${address}:${port}/news/${ticker}/5`;
  request(fetch_address, (error, response, body) => {
    if (!error && response.statusCode === 200) {
      res.send(body);
    }
  })
}

module.exports = {
  get_news
}
