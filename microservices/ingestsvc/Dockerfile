# syntax=docker/dockerfile:1
FROM golang:1.18-alpine As builder
WORKDIR /app
RUN apk add build-base
RUN apk add --no-cache git
COPY . .
RUN mv .netrc /root
ENV GOPRIVATE github.com/frankcai/*
RUN go build -o /bin/ingestsvc

FROM alpine:latest
RUN apk --no-cache add openjdk11
COPY --from=builder /bin/ingestsvc /bin/ingestsvc
COPY --from=builder /app/tika-server-1.21.jar ./tika-server-1.21.jar
COPY /data ./data
ENTRYPOINT [ "/bin/ingestsvc" ]