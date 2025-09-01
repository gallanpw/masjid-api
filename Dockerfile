# Menggunakan image Go yang lengkap
FROM golang:1.24-alpine

# Atur direktori kerja
WORKDIR /app

# Salin semua file sumber
COPY . .

# Unduh dan bangun aplikasi
RUN go mod download
RUN go build -o main .

# Atur variabel lingkungan untuk port
ENV PORT=8080

# Jalankan aplikasi
CMD ["./main"]