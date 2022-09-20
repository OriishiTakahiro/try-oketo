FROM golang:1.18 as builder

WORKDIR /go/src/app
COPY . .

RUN set -x && \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:latest

EXPOSE 80

COPY --from=builder /go/src/app/app /app

CMD ["/app"]
