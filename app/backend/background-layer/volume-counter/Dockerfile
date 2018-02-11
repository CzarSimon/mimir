FROM golang:1.8-alpine3.6 as build
RUN apk update && apk upgrade && apk add --no-cache git nodejs

# Create source directory and boundle source
WORKDIR /usr/src/volume-counter
COPY . /usr/src/volume-counter

# Install go dependnecies
RUN go get github.com/lib/pq
RUN go get github.com/jasonlvhit/gocron
RUN go get github.com/CzarSimon/util
RUN go get github.com/CzarSimon/go-endpoint
RUN go get github.com/CzarSimon/go-file-heartbeat/heartbeat

# Build application
RUN export GOPATH=$GOPATH:$PWD
RUN go build

FROM czarsimon/golang-timezone:1.8-alpine3.6 as run
WORKDIR /root/
COPY --from=build /usr/src/volume-counter/volume-counter .
CMD ["./volume-counter"]
