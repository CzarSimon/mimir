FROM golang:1.8-alpine
RUN apk update && apk upgrade && apk add --no-cache git nodejs
ARG GITHUB_TOKEN
RUN git config --global \
    url."https://$GITHUB_TOKEN:x-oauth-basic@github.com/".insteadOf \
    "https://github.com/"

RUN git config --global --list
