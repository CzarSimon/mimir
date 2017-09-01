FROM golang:1.8-alpine
RUN apk update && apk upgrade && apk add --no-cache git nodejs

RUN mkdir -p /usr/src/counter
COPY . /usr/src/counter

WORKDIR /usr/src/counter

RUN go get github.com/lib/pq
RUN go get github.com/jasonlvhit/gocron
RUN go get github.com/CzarSimon/util
RUN export GOPATH=$GOPATH:$PWD
RUN go build

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /usr/src/counter/counter .
CMD ["./counter"]
