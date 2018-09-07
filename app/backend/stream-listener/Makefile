export DB_HOST=localhost
export DB_PORT=5432
export DB_NAME=streamlistener
export DB_USERNAME=streamlistener
export DB_PASSWORD=password

export SPAM_FILTER_URL=http://localhost:3000
export NEWS_RANKER_URL=http://localhost:5000

run:
	sh start.sh

new-migration:
	alembic revision

db-upgrade:
	alembic upgrade head

db-downgrade:
	alembic downgrade -1
