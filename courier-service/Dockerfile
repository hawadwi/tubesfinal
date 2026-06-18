FROM golang:1.26-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o courier-service .

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/courier-service .

EXPOSE 8081

CMD ["./courier-service"]