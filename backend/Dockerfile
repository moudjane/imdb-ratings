FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY . .

WORKDIR /app/backend


RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main main.go

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/frontend ./frontend

EXPOSE 8080
CMD ["./main"]
