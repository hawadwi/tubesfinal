FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o report-service .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/report-service .

EXPOSE 8083

CMD ["./report-service"]