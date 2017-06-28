'use strict';

const spawn = require('child_process').spawn;
const path = require('path');
const db = require('./database');
const config = require('../config');
const _ = require('lodash');

const rankArticle = (articleInfo, referenceScore, conn, fullStoredArticle = {}) => {
  const storedArticle = stripStoredArticle(fullStoredArticle);
  const { command, path: scriptPath, name } = config.rank_script;
  const scriptFile = path.resolve(scriptPath, name)
  const script = spawn(command, [
    scriptFile,
    JSON.stringify({ articleInfo, referenceScore, storedArticle })
  ]);
  //console.log(JSON.stringify({ articleInfo, referenceScore, storedArticle }));
  _handleScriptData(script);
}

const _handleScriptData = (script) => {
  const terminateProc = childProcces => {
    childProcces.kill('SIGTERM');
  }

  script.stdout.on('data', data => {
    if (data.toString() === 'done') {
      terminateProc(script);
    } else {
      console.log(`scraper stout: ${data}`);
    }
  });

  script.stderr.on('data', data => {
    console.log(`scraper stderr: ${data}`);
    terminateProc(script);
  });
}


const stripStoredArticle = storedArticle => ((_.size(storedArticle) === 0)
    ? storedArticle
    : {
        subject_score: storedArticle.subject_score,
        twitter_references: storedArticle.twitter_references
      }
)


module.exports = {
  rank_article: rankArticle
}
