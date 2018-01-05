FROM python:2

# Create app directory
RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

# Bundle app source
COPY migration.py /usr/src/app/

# install dependencies
RUN pip install rethinkdb==2.3.0.post6
RUN pip install psycopg2==2.6.1

CMD [ "python", "migration.py" ]
