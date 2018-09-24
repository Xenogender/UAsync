
FROM golang:1.20.0-alpine

LABEL maintainer="leoff00"

RUN apk update && apk add --no-cache bash