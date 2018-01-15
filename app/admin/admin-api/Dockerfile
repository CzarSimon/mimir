FROM czarsimon/golang:1.8 as build

# Create source directory
RUN mkdir -p /usr/src/admin-api
WORKDIR /usr/src/admin-api
COPY . /usr/src/admin-api

# Install dependencies
RUN go get github.com/lib/pq
RUN go get github.com/CzarSimon/util
RUN go get github.com/CzarSimon/go-endpoint
RUN go get github.com/CzarSimon/httputil
RUN go get github.com/CzarSimon/mimir/app/lib/go/schema/stock
RUN go get github.com/CzarSimon/mimir/app/lib/go/schema/spam

# Build application
RUN go build

FROM alpine:latest
WORKDIR /root/
COPY --from=build /usr/src/admin-api/admin-api .
CMD ["./admin-api"]
