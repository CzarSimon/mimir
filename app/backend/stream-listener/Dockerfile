FROM python:3.6-alpine

WORKDIR /src/stream-listener
COPY . .

RUN pip install -r requirements.txt

CMD ["sh", "start.sh"]
