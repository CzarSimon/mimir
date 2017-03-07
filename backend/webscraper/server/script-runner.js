'use strict';

const spawn = require('child_process').spawn
    , path = require('path')
    , db = require('./database')
    , config = require('../config')
    , _ = require('lodash');

const rank_article = (article_info, reference_score, conn, stored_article = {}) => {
  const { command, path: script_path, name } = config.rank_script;
  const script_file = path.resolve(script_path, name)

  const script = spawn(command, [
    script_file,
    JSON.stringify({
      article_info: article_info,
      reference_score: reference_score
    })
  ]);
  _handle_script_data(script, conn, reference_score, stored_article);
}

const _handle_script_data = (script, conn, ref_score, stored_article) => {
  const _terminate_proc = (child_proc) => {
    child_proc.kill('SIGTERM');
  }

  script.stdout.on('data', data => {
    // console.log(`stout: ${data}`);
    try {
      const new_article = JSON.parse(data);
      if ((_.size(stored_article) === 0)) {
        db.insert_articles(new_article, conn);
      } else {
        const new_subj_scores = Object.assign({}, stored_article.subject_score, new_article.subject_score)
            , compound_score = _.mapValues(new_subj_scores, subj_score => subj_score + ref_score)
            , twitter_references = _.union(stored_article.twitter_references, new_article.twitter_references)
            , updated_article = Object.assign({}, new_article,
              {
                subject_score: new_subj_scores,
                compound_score: compound_score,
                twitter_references: twitter_references
              }
            );
        //console.log(updated_article);
        db.update_article(new_article.id, updated_article, conn);
      }
    } catch (e) {
      console.log(" ------- There was an error --------- ");
      console.log(e);
      //console.log("Data: ", data.toString());
    } finally {
      _terminate_proc(script);
    }
  });

  script.stderr.on('data', data => {
    console.log(`stderr: ${data}`);
    _terminate_proc(script);
  });
}

module.exports = {
  rank_article: rank_article
}
