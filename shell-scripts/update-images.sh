function update() {
  local image=$1
  echo "Updating $image"
  docker pull $image
}

update "postgres"
update "rethinkdb"
update "czarsimon/spam-filter"
update "czarsimon/mimir-app-server"
update "czarsimon/news-server"
update "czarsimon/stream-listener"
update "czarsimon/news-ranker"
update "czarsimon/volume-counter"
update "czarsimon/mean-stdev-calc"
update "czarsimon/news-clusterer"
update "czarsimon/mimir-admin-page"
