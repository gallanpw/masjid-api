# Tahap 1: Builder
# Gunakan image Go yang lengkap untuk mengompilasi aplikasi
FROM golang:1.24-alpine AS builder

# Atur direktori kerja
WORKDIR /app

# Salin file mod dan sum
COPY go.mod go.sum ./

# Unduh dependensi
RUN go mod download

# Salin semua file sumber
COPY . .

# Bangun aplikasi ke file biner bernama 'main'
# CGO_ENABLED=0 untuk membuat biner yang berdiri sendiri
RUN CGO_ENABLED=0 go build -o main .

# ---

# Tahap 2: Tahap Akhir (Runtime)
# Salin file biner dari tahap builder ke tahap akhir
COPY --from=builder /app/main .

# Jalankan aplikasi. Ini hanya menjalankan biner, tanpa 'go'
CMD ["./main"]