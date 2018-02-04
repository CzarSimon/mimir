go build

# Tweet db config
export DB_HOST=mimir-dev.news
export DB_PORT=32201
export DB_USER=simon
export DB_PASSWORD=$PG_PASSWORD
export DB_NAME=mimirprod

# Tweet handler config
export SEARCH_SERVER_PORT=7000

./search-server
