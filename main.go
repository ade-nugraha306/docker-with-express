package main

import (
    "log"

    "go-relog/config"
    "go-relog/handlers"

    "github.com/gin-gonic/gin" // Import Gin
)

func main() {
    config.ConnectDB()

    r := gin.Default() // Inisialisasi Gin router

    // Definisikan routes
    r.POST("/register", handlers.Register)
    r.POST("/login", handlers.Login)

    log.Println("Server running on port 8080 with Gin!")
    r.Run(":8080") // Jalankan server Gin
}