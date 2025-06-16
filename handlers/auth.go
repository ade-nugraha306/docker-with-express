package handlers

import(
	"context"
	"net/http"
	"time"

	"go-relog/config"
	"go-relog/models"
	"go-relog/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Register (c*gin.Context){
	var user models.User
	if err := c.ShouldBindBodyWithJSON(&user); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid req body!"})
		return
	}

	hashedPassword, err := utils.HashPw(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash!"})
		return
	}
	user.Password = hashedPassword

	collection := config.GetCollection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, user)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func Login(c *gin.Context) {
	var loginRequest struct { // Buat struct sementara untuk request login
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
		return
	}

	collection := config.GetCollection("user") // Pastikan ini "user"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result models.User
	// Cari user berdasarkan email
	err := collection.FindOne(ctx, bson.M{"email": loginRequest.Email}).Decode(&result)
	if err != nil {
		// Jika email tidak ditemukan
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials. Email tidak terdaftar!"})
		return
	}

	// Cek password
	if !utils.CheckHashPw(loginRequest.Password, result.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful!"})
}