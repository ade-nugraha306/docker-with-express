# Base image untuk Go
FROM golang:1.24-alpine

# Set direktori kerja
WORKDIR /app

# Copy go.mod dan go.sum
COPY go.mod go.sum ./

# Download dependensi
RUN go mod download

# Copy kode aplikasi
COPY . .

# Build aplikasi
RUN go build -o main .

# Expose port
EXPOSE 8080

# Command untuk menjalankan aplikasi
CMD go run main.go