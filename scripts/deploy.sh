sctl unlock
sctl check service ls

services=(
  "app-db"
  "tweet-db"
  "news-db"
  "price-db"
  "news-clusterer"
  "spam-filter"
  "news-ranker"
  #"volume-counter"
  "price-service"
  "app-server"
  "news-server"
  "search-service"
  "price-server"
  "stream-listener"
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
