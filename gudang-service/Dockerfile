FROM golang:1.26-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o gudang-service .

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/gudang-service .

EXPOSE 8080

CMD ["./gudang-service"]