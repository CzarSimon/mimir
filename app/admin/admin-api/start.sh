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
export ADMIN_API_ACCESS_KEY=0b6994adf3012f037b84d4362c65d70ad278848e00963b6502b04793742464d7a2beea77cd120512e4e7a318af3948c6fe647b235d520ce310ffcec8a6aa686a
export TOKEN_EXPIRIY_MINUTES=10

./admin-api
