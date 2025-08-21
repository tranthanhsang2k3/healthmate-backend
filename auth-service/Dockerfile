# Dockerfile - user-service

FROM golang:1.23 AS builder

WORKDIR /app

COPY auth-service/go.mod auth-service/go.sum ./
RUN go mod download

COPY auth-service/. ./

RUN go build -o auth-service .

FROM gcr.io/distroless/base-debian11

COPY --from=builder /app/auth-service .

EXPOSE 9000

CMD ["./auth-service"]
