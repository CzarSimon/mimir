FROM czarsimon/golang:1.8 as build

RUN mkdir -p /usr/src/app-server
COPY . /usr/src/app-server

WORKDIR /usr/src/app-server

RUN go get github.com/lib/pq
RUN go get github.com/CzarSimon/util
RUN go get github.com/CzarSimon/go-endpoint
RUN go get github.com/CzarSimon/httputil
RUN go get github.com/CzarSimon/mimir/app/lib/go/schema
RUN export GOPATH=$GOPATH:$PWD
RUN go build

FROM alpine:latest
WORKDIR /root/
COPY --from=build /usr/src/app-server/app-server .
EXPOSE 3000
CMD ["./app-server"]
