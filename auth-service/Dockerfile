# Dockerfile - user-service

FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o auth-service .

FROM gcr.io/distroless/base-debian11

COPY --from=builder /app/auth-service .

EXPOSE 9000

CMD ["./auth-service"]
