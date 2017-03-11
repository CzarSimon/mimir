'use strict';

const request = require('request')
const _ = require('lodash')
const { clusterServer } = require('../config')


const pickClusterAttributes = article => (
  _.pick(article, ['id', 'title', 'timestamp', 'reference_score', 'subject_score'])
)


const sendToClustering = articleInfo => {
  const { id, title, timestamp, reference_score, subject_score } = articleInfo
  const info = {
    urlHash: id,
    title: title,
    date: timestamp
  }
  _.forEach(subject_score, (subjectScore, ticker) => {
    _sendRequest(_formatArticleInfo(info, ticker, subjectScore, reference_score))
  })
}


const _formatArticleInfo = (articleInfo, ticker, subjectScore, referenceScore) => {
  return Object.assign({}, articleInfo, {
    ticker: ticker,
    score: {
      subjectScore: subjectScore,
      referenceScore: referenceScore
    }
  })
}

const _sendRequest = articleInfo => {
  const { protocol, host, port } = clusterServer
  const endpoint = protocol + "://" + host + ":" + port + "/api/cluster-article";
  request.post({
    url: endpoint,
    json: true,
    headers: {
      "Content-Type": "application/json"
    },
    body: articleInfo
  }, (err, res, body) => {
    if (err) {
      console.log("Error sending to cluster", err);
      console.log("Body: ", JSON.parse(body));
    }
  })
}

module.exports = {
  sendToClustering: sendToClustering,
  pickClusterAttributes: pickClusterAttributes
}
