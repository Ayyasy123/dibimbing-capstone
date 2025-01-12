# Menggunakan Golang 1.20 dengan Alphine linux sebagai base image
FROM golang:1.23-alpine

# Set working directory di dalam container
WORKDIR /app

# Copy go.mod dan go.user terlebih dahulu untuk menghindari rebuild dependensi saat kode berubah
COPY go.mod go.sum ./

# Download dan verifikasi dependensi Go
RUN go mod download && go mod tidy

# Copy seluruh kode aplikasi ke dalam container
COPY . .

# Build aplikasi go
RUN go build -o main .

# Expose port 8080 agar aplikasi bisa diakses dari luar container
EXPOSE 8080

# menjalankan aplikasi go
CMD [ "./main" ]



