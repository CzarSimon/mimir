FROM czarsimon/golang:1.8 as build

# Create source folder
WORKDIR /usr/src/tweet-handler
COPY . .

# Install dependencies
RUN go get github.com/lib/pq
RUN go get github.com/CzarSimon/mimir/app/lib/go/schema
RUN go get github.com/CzarSimon/httputil
RUN go get github.com/CzarSimon/go-endpoint
RUN go get github.com/CzarSimon/util
RUN go get github.com/goware/urlx

# Build
RUN export GOPATH=$GOPATH:$PWD
RUN go build
RUN go test

FROM alpine:latest as run
WORKDIR /root/
COPY --from=build /usr/src/tweet-handler/tweet-handler .
CMD ["./tweet-handler"]
