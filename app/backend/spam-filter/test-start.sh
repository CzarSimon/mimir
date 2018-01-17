export PG_NAME=mimirprod
export PG_USER=simon
export PG_HOST=mimir-dev.news
export PG_PASSWORD=$PG_PASSWORD
export PG_PORT=32201

export CASHTAG_THRESHOLD=0.8

gunicorn -w 1 -b 127.0.0.1:21000 main
