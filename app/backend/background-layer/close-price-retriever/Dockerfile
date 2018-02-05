FROM czarsimon/golang:1.8 as build

# Create source directory and bundle app source
WORKDIR /usr/src/close-price-retriever
COPY . .

# Install go dependecies
RUN go get github.com/lib/pq
RUN go get github.com/CzarSimon/util
RUN go get github.com/CzarSimon/go-endpoint
RUN go get github.com/CzarSimon/mimir/app/lib/go/api
RUN go get github.com/CzarSimon/mimir/app/lib/go/schema/stock

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
COPY --from=build /usr/src/close-price-retriever/close-price-retriever .
CMD ["./close-price-retriever"]
