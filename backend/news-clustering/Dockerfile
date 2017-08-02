FROM golang:1.8-alpine
RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir -p /usr/src/news-clusterer
COPY . /usr/src/news-clusterer

WORKDIR /usr/src/news-clusterer

RUN go get github.com/lib/pq
RUN go get github.com/CzarSimon/util
RUN export GOPATH=$GOPATH:$PWD
RUN go build

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /usr/src/news-clusterer/news-clusterer .

EXPOSE 6000
CMD ["./news-clusterer"]
