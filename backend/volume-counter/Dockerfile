FROM golang:1.8-alpine
RUN apk update && apk upgrade && apk add --no-cache git nodejs

RUN mkdir -p /usr/src/counter
COPY . /usr/src/counter

WORKDIR /usr/src/counter

RUN go get github.com/lib/pq
RUN go get github.com/jasonlvhit/gocron
RUN export GOPATH=$GOPATH:$PWD
RUN go build

CMD ["./counter"]
