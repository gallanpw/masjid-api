# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux
RUN go build -o main .

# Final stage
FROM alpine:3.18

WORKDIR /

ENV PORT=8080

COPY --from=builder /app/main .

RUN chmod +x main

CMD ["./main"]
