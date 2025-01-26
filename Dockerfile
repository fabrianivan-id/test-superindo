# Build Stage
FROM golang:1.23.5 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./*.go ./
RUN go build -o main .

# Final Stage
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/main .
COPY /migrations /migrations

EXPOSE 8080

CMD ["./main"]
