# syntax=docker/dockerfile:1
FROM golang:1.18-alpine As builder
RUN apk add build-base
WORKDIR /app
COPY . .
RUN go build -o monolith

FROM alpine:latest
WORKDIR /app
RUN apk --no-cache add openjdk11
COPY --from=builder /app/monolith ./
COPY --from=builder /app/tika-server-1.21.jar ./
COPY /data ./data
ENTRYPOINT [ "./monolith" ]