FROM golang:1.8-alpine
RUN apk update && apk upgrade && apk add --no-cache git nodejs
RUN git config --global \
    url."https://e0f04d042777d83c8084fcd75126e2fe8573aed8:x-oauth-basic@github.com/".insteadOf \
    "https://github.com/"

RUN mkdir -p /usr/src/price-server
COPY . /usr/src/price-server

WORKDIR /usr/src/price-server

RUN echo $GITHUB_TOKEN
RUN go get github.com/lib/pq
RUN go get github.com/CzarSimon/util
RUN go get github.com/CzarSimon/mimir/lib/stock
RUN export GOPATH=$GOPATH:$PWD
RUN go build

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /usr/src/price-server/price-server .
CMD ["./price-server"]
