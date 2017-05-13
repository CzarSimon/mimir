FROM golang:1.8-alpine
RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir -p /usr/src/news-server
COPY . /usr/src/news-server

WORKDIR /usr/src/news-server

RUN go get gopkg.in/gorethink/gorethink.v2
RUN go get github.com/julienschmidt/httprouter
RUN export GOPATH=$GOPATH:$PWD
RUN go build

EXPOSE 5050
CMD ["./news-server"]
