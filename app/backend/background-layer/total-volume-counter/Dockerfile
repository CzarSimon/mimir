FROM golang:1.8-alpine3.6 as build
RUN apk update && apk upgrade && apk add --no-cache git

# Create source directory and bundle app source
WORKDIR /usr/src/total-volume-counter
COPY . .

# Install go dependecies
RUN go get github.com/lib/pq
RUN go get github.com/CzarSimon/util
RUN go get github.com/CzarSimon/go-endpoint

# Build app
RUN export GOPATH=$GOPATH:$PWD
RUN go build

FROM czarsimon/golang-timezone:1.8-alpine3.6 as run
WORKDIR /root/
COPY --from=build /usr/src/total-volume-counter/total-volume-counter .
CMD ["./total-volume-counter"]
