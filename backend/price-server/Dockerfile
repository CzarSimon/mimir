FROM golang:1.8-alpine
RUN apk update && apk upgrade && apk add --no-cache git nodejs

RUN mkdir -p /usr/src/price-server
COPY . /usr/src/price-server

WORKDIR /usr/src/price-server

RUN go get github.com/lib/pq
RUN go get github.com/CzarSimon/util
RUN export GOPATH=$GOPATH:$PWD
RUN go build

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /usr/src/price-server/price-server .
CMD ["./price-server"]
