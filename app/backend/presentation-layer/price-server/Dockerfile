FROM czarsimon/golang:1.8 as build

# Create application folder
RUN mkdir -p /usr/src/price-server
COPY . /usr/src/price-server
WORKDIR /usr/src/price-server

# Download dependencies
RUN go get github.com/lib/pq
RUN go get github.com/CzarSimon/util
RUN go get github.com/CzarSimon/httputil
RUN go get github.com/CzarSimon/go-endpoint
RUN go get github.com/CzarSimon/mimir/app/lib/go/schema/stock

# Build
RUN export GOPATH=$GOPATH:$PWD
RUN go build

FROM alpine:latest as run
WORKDIR /root/
COPY --from=build /usr/src/price-server/price-server .
CMD ["./price-server"]
