go build

# Tweet db config
export DB_HOST=mimir-dev.news
export DB_PORT=32201
export DB_USER=simon
export DB_PASSWORD=$PG_PASSWORD
export DB_NAME=mimirprod

# Spam filter config
export SPAM_FILTER_HOST=mimir-dev.news
export SPAM_FILTER_PROTOCOL=http
export SPAM_FILTER_PORT=30073

# News ranker config
export NEWS_RANKER_HOST=mimir-dev.news
export NEWS_RANKER_PROTOCOL=http
export NEWS_RANKER_PORT=30288

# Tweet handler config
export TWEET_HANDLER_PORT=2000

./tweet-handler
