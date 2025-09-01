# Menggunakan image Go yang lengkap
FROM golang:1.24-alpine
# Atur direktori kerja
WORKDIR /app
# Salin go.mod dan go.sum untuk mengunduh dependensi
COPY go.mod go.sum ./
# Unduh semua dependensi
RUN go mod download
# Salin semua file sumber
COPY . .
# Bangun (compile) aplikasi ke file biner bernama 'main'
RUN go build -o main .
# Mengekspos port 8080 agar dapat diakses oleh Railway
EXPOSE 8080
# Jalankan aplikasi
CMD ["./main"]