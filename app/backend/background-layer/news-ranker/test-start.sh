export DB_NAME=mimirprod
export DB_USER=simon
export DB_HOST=mimir-dev.news
export DB_PASSWORD=$PG_PASSWORD
export DB_PORT=32197

export NEWS_CLUSTERER_PROTOCOL=http
export NEWS_CLUSTERER_HOST=mimir-dev.news
export NEWS_CLUSTERER_PORT=32459

export NEWS_RANKER_PORT=5000

echo "Building service"
go build
echo "Build done"
./news-ranker
