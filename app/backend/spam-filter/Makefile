export DB_HOST=localhost
export DB_PORT=5432
export DB_NAME=spamfilter
export DB_USERNAME=spamfilter
export DB_PASSWORD=password
export CASHTAG_THRESHOLD=0.8
export TRAIN_MODEL=FALSE

run:
	sh start.sh

db-init:
	rm -rf migrations
	flask db init

db-migrate:
	flask db migrate

db-upgrade:
	flask db upgrade

db-downgrade:
	flask db downgrade
