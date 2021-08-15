FROM golang:1.16.7-alpine3.14 AS builder
COPY . /app
WORKDIR /app
RUN mkdir ./build && \
    go build -o ./build ./...

FROM alpine:3.14.1
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder app/build/covid19-india /app/
COPY --from=builder app/docs/ /app/docs
WORKDIR /app
CMD ["./covid19-india"]
