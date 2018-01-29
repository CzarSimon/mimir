FROM czarsimon/golang:1.8 as build
RUN apk update && apk upgrade && apk add --no-cache git

RUN mkdir -p /usr/src/news-server
COPY . /usr/src/news-server
WORKDIR /usr/src/news-server

RUN go get github.com/lib/pq
RUN go get github.com/CzarSimon/util
RUN go get github.com/CzarSimon/httputil
RUN go get github.com/CzarSimon/go-endpoint
RUN go get github.com/julienschmidt/httprouter

RUN export GOPATH=$GOPATH:$PWD
RUN go build

FROM alpine:latest as run
WORKDIR /root/
COPY --from=build /usr/src/news-server/news-server .
EXPOSE 5050
CMD ["./news-server"]
