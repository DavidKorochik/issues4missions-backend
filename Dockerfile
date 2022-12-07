# Build the Golang app
FROM golang:1.19.4-alpine3.16 AS builder

WORKDIR /app

COPY . .

RUN go build -o main cmd/main/main.go

# Run the exe file
FROM alpine:3.17

WORKDIR /app

COPY --from=builder /app/main .

COPY server.env .

EXPOSE 5000

CMD ["/app/main"]
