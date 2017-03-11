FROM golang:1.8

RUN mkdir -p /usr/src/news-clusterer
COPY . /usr/src/news-clusterer

WORKDIR /usr/src/news-clusterer

RUN go get gopkg.in/gorethink/gorethink.v2
RUN export GOPATH=$GOPATH:$PWD
RUN go build

EXPOSE 6000
CMD ["./news-clusterer"]
