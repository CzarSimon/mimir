FROM czarsimon/golang:1.8 as build

# Create source directory and bundle app source
WORKDIR /usr/src/price-retriever
COPY . .

# Install go dependecies
RUN go get github.com/lib/pq
RUN go get github.com/jasonlvhit/gocron
RUN go get github.com/CzarSimon/go-endpoint
RUN go get github.com/CzarSimon/go-file-heartbeat/heartbeat
RUN go get github.com/CzarSimon/mimir/app/lib/go/api
RUN go get github.com/CzarSimon/mimir/app/lib/go/schema/stock
RUN go get github.com/CzarSimon/util

# Build app
RUN export GOPATH=$GOPATH:$PWD
RUN go build

FROM alpine:latest as run

# Installing and updating root certs
RUN apk update && \
    apk upgrade && \
    apk add --no-cache ca-certificates && \
    update-ca-certificates

WORKDIR /root/
COPY --from=build /usr/src/price-retriever/price-retriever .
CMD ["./price-retriever"]
