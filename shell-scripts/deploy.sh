sctl unlock
sctl check service ls

services=(
  "app-rdb"
  "tweet-db"
  "news-db"
  "price-db"
  "price-service"
  "news-clusterer"
  "news-ranker"
  "volume-counter"
  "spam-filter"
  #"stream-listener"
  "app-server"
  "news-server"
  "search-service"
  "price-server"
)

start_service() {
  service=$1
  sctl start $service
  sleep 20
}

for service in "${services[@]}"
do
  start_service $service
done

sctl check service ls
sctl lock
