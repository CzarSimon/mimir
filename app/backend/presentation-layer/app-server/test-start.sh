export DB_NAME=mimirprod
export DB_USER=simon
export DB_HOST=mimir-dev.news
export DB_PASSWORD=$PG_PASSWORD
export DB_PORT=30012

export APP_SERVER_PORT=3000

go build
./app-server
