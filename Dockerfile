# Tahap 1: Builder
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o main .

# ---

# Tahap 2: Tahap Akhir (Runtime)
FROM alpine:3.18

# Buat direktori aplikasi dan atur sebagai direktori kerja
WORKDIR /app

# Salin file biner dari tahap builder ke direktori /app
COPY --from=builder /app/main /app/main

# Jalankan aplikasi. Ini secara eksplisit menjalankan biner dari /app
CMD ["./main"]