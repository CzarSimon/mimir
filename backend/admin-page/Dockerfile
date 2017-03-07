FROM golang:1.8-alpine
RUN apk update && apk upgrade && apk add --no-cache git nodejs

RUN mkdir -p /usr/src/app
COPY . /usr/src/app

WORKDIR /usr/src/app
RUN npm install
RUN npm run build

WORKDIR /usr/src/app/server

RUN go get github.com/lib/pq
RUN go get gopkg.in/gorethink/gorethink.v2
RUN export GOPATH=$GOPATH:$PWD
RUN go build

EXPOSE 8000
CMD ["./server"]
