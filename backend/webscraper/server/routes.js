'use strict';

const db = require('./database')
    , _ = require('lodash')
    , moment = require('moment')
    , script_runner = require('./script-runner')
    , { twitter_users, reference_weight } = require('../config').rank_script;

const get_news_articles = (request, result, conn) => {
  const { ticker, top } = request.params;
  db.fetchTopArticles(ticker, parseInt(top), moment.utc().format('YYYY-MM-DD'), conn, (err, res) => {
    result.send(res);
  });
}

const rank_articles = (request, result, conn) => {
  const articles_info = request.body;
  if (articles_info) {
    result.send({ success: true });
    for (let url of articles_info.urls) {
      const article_info = Object.assign({}, articles_info, { url: url });
      _check_for_new_article(article_info, conn);
    }
  } else {
    result.send({ success: false })
  }
}

//This function needs error handling in ca
const _check_for_new_article = (article_info, conn) => {
  db.check_for_article(article_info.url, conn, (err, res) => {
    if (res.length > 0) {
      //Article already exist
      const stored_article = res[0];
      if (_new_ticker_reference(stored_article, article_info)) {
        //New ticker in twitter reference
        //console.log("New ticker in twitter reference");
        _rank_with_new_ticker(stored_article, article_info, conn);
      } else {
        //Same tickers in twitter reference
        //console.log("Same tickers in twitter reference");
        _update_ref_score(stored_article, article_info, conn);
      }
    } else {
      //New article
      //console.log("New article");
      _rank_new_article(article_info, conn);
    }
  });
}

const _rank_with_new_ticker = (stored_article, new_article_info, conn) => {
  const updated_ref_score = _calc_reference_score(new_article_info.author.follower_count, stored_article.reference_score);
  script_runner.rank_article(new_article_info, updated_ref_score, conn, stored_article);
}

const _rank_new_article = (article_info, conn) => {
  const ref_score = _calc_reference_score(article_info.author.follower_count);
  script_runner.rank_article(article_info, ref_score, conn);
}

const _update_ref_score = (storedArticle, new_article_info, conn) => {
  const { subject_score, id, reference_score, twitter_references } = storedArticle;
  const { author, subjects } = new_article_info;
  if (!twitter_references.includes(author.id)) {
    //console.log("New user posted");
    const updated_ref_score = _calc_reference_score(author.follower_count, reference_score);
    db.update_article(id,
      Object.assign({}, _.pick(storedArticle, ['id', 'title', 'timestamp', 'subject_score']),
      {
        reference_score: updated_ref_score,
        compound_score: _calc_new_compound_score(subject_score, updated_ref_score),
        twitter_references: _.concat(twitter_references, author.id)
      }
    ), conn);
  }
}

const _new_ticker_reference = (stored_article, new_article_info) => {
  const stored_tickers = _.keys(stored_article.subject_score);
  const new_tickers = _.map(new_article_info.subjects, subject => subject.ticker);
  return (_.difference(_.union(new_tickers, stored_tickers), stored_tickers).length > 0);
}

const _calc_new_compound_score = (subject_scores, ref_score) => {
  return _.mapValues(subject_scores, subj_score => subj_score + ref_score);
}

const _calc_reference_score = (followers, prev_score = 0.0) => {
  const additional_score = reference_weight * parseFloat(followers) / twitter_users;
  return additional_score + prev_score;
}

module.exports = {
  get_news_articles: get_news_articles,
  rank_articles: rank_articles
};
