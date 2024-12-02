FROM golang:1.23.2 AS builder

WORKDIR /build

COPY . .

RUN ls

RUN cd auth && \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o /server

FROM golang:1.23.2

WORKDIR /app

COPY --from=builder /server server
COPY --from=builder /build/server/migrations migrations

CMD ["/app/server"]