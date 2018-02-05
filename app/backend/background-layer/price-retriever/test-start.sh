export PRICE_DB_NAME=mimirprod
export PRICE_DB_USER=simon
export PRICE_DB_HOST=mimir-dev.news
export PRICE_DB_PASSWORD=$PG_PASSWORD
export PRICE_DB_PORT=30533

export TICKER_DB_NAME=mimirprod
export TICKER_DB_USER=simon
export TICKER_DB_HOST=mimir-dev.news
export TICKER_DB_PASSWORD=$PG_PASSWORD
export TICKER_DB_PORT=30012

export TIMEZONE="America/New_York"

export HEARTBEAT_FILE="$PWD/heartbeat.txt"
export HEARTBEAT_INTERVAL=15

export EXCHANGE_OPEN_HOUR=9
export EXCHANGE_CLOSE_HOUR=16

echo "Building service"
go build
echo "Build done"
./price-retriever
