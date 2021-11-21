FROM golang:1.16-buster AS builder

WORKDIR /app
COPY . /app

RUN CGO_ENABLED=0 go build  -mod=vendor -o bin/service ./cmd/main.go


FROM debian:buster
RUN apt-get update \
    && apt-get install -y ca-certificates \
    && update-ca-certificates

COPY --from=builder /app/bin/service /service

COPY migrations /migrations
COPY templates /templates

CMD ["/service"]
