export DB_NAME=mimirprod
export DB_USER=simon
export DB_HOST=mimir-dev.news
export DB_PASSWORD=$PG_PASSWORD
export DB_PORT=32197

export NEWS_CLUSTERER_PORT=6000

echo "Building service"
go build
echo "Build done"
./news-clusterer
