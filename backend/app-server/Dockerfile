FROM golang:1.8-alpine
RUN apk update && apk upgrade && apk add --no-cache git nodejs

RUN mkdir -p /usr/src/app-server
COPY . /usr/src/app-server

WORKDIR /usr/src/app-server

RUN go get github.com/lib/pq
RUN go get github.com/CzarSimon/util
RUN export GOPATH=$GOPATH:$PWD
RUN go build

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /usr/src/app-server/app-server .
EXPOSE 3000
CMD ["./app-server"]
