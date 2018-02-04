FROM czarsimon/golang:1.8 as build

# Create source folder
RUN mkdir -p /usr/src/tweet-handler
COPY . /usr/src/tweet-handler
WORKDIR /usr/src/tweet-handler

# Install dependencies
RUN go get github.com/lib/pq
RUN go get github.com/CzarSimon/mimir/app/lib/go/schema
RUN go get github.com/CzarSimon/httputil
RUN go get github.com/CzarSimon/go-endpoint
RUN go get github.com/CzarSimon/util

# Build
RUN export GOPATH=$GOPATH:$PWD
RUN go build

FROM alpine:latest as run
WORKDIR /root/
COPY --from=build /usr/src/tweet-handler/tweet-handler .
CMD ["./tweet-handler"]
