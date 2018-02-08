export DB_NAME=mimirprod
export DB_USER=simon
export DB_HOST=mimir-dev.news
export DB_PASSWORD=$PG_PASSWORD
export DB_PORT=32201

echo "Building service"
go build
echo "Build done"
./total-volume-counter
