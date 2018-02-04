FROM czarsimon/golang:1.8 as build

# Create source folder and bundle applicaiton source
RUN mkdir -p /usr/src/news-clusterer
COPY . /usr/src/news-clusterer
WORKDIR /usr/src/news-clusterer

# Install dependencies
RUN go get github.com/lib/pq
RUN go get github.com/CzarSimon/util
RUN go get github.com/CzarSimon/httputil
RUN go get github.com/CzarSimon/go-endpoint
RUN export GOPATH=$GOPATH:$PWD
RUN go build

# Package application and preprare run environment
FROM alpine:latest as run
WORKDIR /root/
COPY --from=build /usr/src/news-clusterer/news-clusterer .
EXPOSE 6000
CMD ["./news-clusterer"]
