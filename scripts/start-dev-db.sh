NETWORK=mimir-net
DB_NAME=mimir-db
docker rm -f $DB_NAME

docker run -d --name $DB_NAME \
  -p 5432:5432 --network $NETWORK \
  -e POSTGRES_PASSWORD=password \
  --restart always postgres:10.2-alpine

NAME=mimir-dbui
docker rm -f $NAME

docker run -d --name $NAME \
  -p 8081:8081 --network $NETWORK \
  --restart always sosedoff/pgweb

echo "Installing databases"
sleep 3
docker exec -i $DB_NAME psql -U postgres < dev-db.sql
