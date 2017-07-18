FROM golang:1.8-alpine
RUN apk update && apk upgrade && apk add --no-cache git nodejs

RUN mkdir -p /usr/src/search-recomender
COPY . /usr/src/search-recomender

WORKDIR /usr/src/search-recomender

RUN go get github.com/lib/pq
RUN go get github.com/CzarSimon/util
RUN export GOPATH=$GOPATH:$PWD
RUN go build

EXPOSE 7000
CMD ["./search-recomender"]
