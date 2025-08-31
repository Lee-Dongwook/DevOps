FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o service main.go

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/service .
EXPOSE 8080
CMD ["./service"]
