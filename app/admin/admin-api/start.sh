go build # Builds the api server

# Tweet db config
export TWEET_DB_HOST=mimir-dev.news
export TWEET_DB_PORT=32201
export TWEET_DB_USER=simon
export TWEET_DB_PASSWORD=$PG_PASSWORD
export TWEET_DB_NAME=mimirprod

# App db config
export APP_DB_HOST=mimir-dev.news
export APP_DB_PORT=30012
export APP_DB_USER=simon
export APP_DB_PASSWORD=$PG_PASSWORD
export APP_DB_NAME=mimirprod

# Server address config
export ADMIN_API_PORT=3000

# Auth credentials
export ADMIN_API_ACCESS_KEY=2464ca60e65117abc5656bcef7b019d001c78e20f6df9ad85f7574d890f6f9884504cd39f9d890ce8791d619f02e02b17e4be1018423234ed839355d21adcc0f
export TOKEN_EXPIRIY_MINUTES=1

./admin-api
