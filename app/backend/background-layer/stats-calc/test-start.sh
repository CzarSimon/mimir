export TWEET_DB_NAME=mimirprod
export TWEET_DB_USER=simon
export TWEET_DB_HOST=mimir-dev.news
export TWEET_DB_PASSWORD=$PG_PASSWORD
export TWEET_DB_PORT=32201

export APP_DB_NAME=mimirprod
export APP_DB_USER=simon
export APP_DB_HOST=mimir-dev.news
export APP_DB_PASSWORD=$PG_PASSWORD
export APP_DB_PORT=30012

export STABLE_DATE="2018-01-20"

echo "Building service"
go build
echo "Build done"
./stats-calc
