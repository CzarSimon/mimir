FROM golang:1.8-alpine
RUN apk update && apk upgrade && apk add --no-cache git nodejs

RUN mkdir -p /usr/src/price-service
COPY . /usr/src/price-service

WORKDIR /usr/src/price-service

RUN go get github.com/lib/pq
RUN go get github.com/jasonlvhit/gocron
RUN go get github.com/FlashBoys/go-finance
RUN go get github.com/CzarSimon/util
RUN export GOPATH=$GOPATH:$PWD
RUN go build

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /usr/src/price-service/price-service .
CMD ["./price-service"]
