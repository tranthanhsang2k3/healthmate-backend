FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o auth-service ./cmd

FROM ubuntu:22.04

WORKDIR /app

# Cài netcat + tzdata để chờ DB và xử lý timezone
RUN apt-get update && apt-get install -y netcat tzdata ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/auth-service .
COPY wait-for-it.sh .
RUN chmod +x wait-for-it.sh

EXPOSE 9000

CMD ["./wait-for-it.sh", "postgres", "5432", "--", "./auth-service"]