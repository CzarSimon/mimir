sctl unlock
sctl check service ls

services=(
  "app-server"
  "news-server"
  "search-service"
  "price-server"
  "price-service"
  #"stream-listener"
  "volume-counter"
  "spam-filter"
  "news-ranker"
  "news-clusterer"
  "app-rdb"
  "tweet-db"
  "news-db"
  "price-db"
)

stop_service() {
  service=$1
  sctl stop $service
  sleep 1
}

for service in "${services[@]}"
do
  stop_service $service
done

sctl check service ls
sctl lock
