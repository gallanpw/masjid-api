# Tahap 1: Builder
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Tahap 2: Final image yang ringan
FROM alpine:3.18
WORKDIR /app # Atur direktori kerja ke /app di tahap akhir
COPY --from=builder /app/main .
# Tentukan perintah untuk menjalankan aplikasi
CMD ["./main"]