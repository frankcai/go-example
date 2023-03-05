# syntax=docker/dockerfile:1
FROM golang:1.18-alpine As builder
WORKDIR /app
RUN apk add build-base \
    git
COPY . .
ENV GOPRIVATE github.com/frankcai/*
RUN go mod download
# go test -bench=Bench -benchmem -cpu=1 -benchtime=1x
ENTRYPOINT [ "go", "test", "-bench=Bench", "-benchmem", "-cpu=1", "-benchtime=1x"]