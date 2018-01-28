FROM golang:1.8-alpine as build
RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir -p /usr/src/search-service
COPY . /usr/src/search-service
WORKDIR /usr/src/search-service

RUN go get github.com/lib/pq
RUN go get github.com/CzarSimon/httputil
RUN go get github.com/CzarSimon/go-endpoint
RUN go get github.com/CzarSimon/util

RUN export GOPATH=$GOPATH:$PWD
RUN go build

FROM alpine:latest as run
WORKDIR /root/
COPY --from=build /usr/src/search-service/search-service .
EXPOSE 7000
CMD ["./search-service"]
