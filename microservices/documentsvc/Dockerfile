# syntax=docker/dockerfile:1
FROM golang:1.18-alpine As builder
WORKDIR /app
RUN apk add build-base
RUN apk add --no-cache git
COPY . .
RUN mv .netrc /root
ENV GOPRIVATE github.com/frankcai/*
RUN go mod download
RUN go build -o /bin/documentsvc

FROM alpine:latest
COPY --from=builder /bin/documentsvc /bin/documentsvc
ENTRYPOINT [ "/bin/documentsvc" ]