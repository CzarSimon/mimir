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

echo "Building service"
go build
echo "Build done"
./close-price-retriever
