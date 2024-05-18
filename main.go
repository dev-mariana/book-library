package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Book struct {
	id string `json:"id" gorm:"primary_key"`  
	title string `json:"title"`
	description string `json:"description"`
	author string 	`json:"author"`
	genre string 	`json:"genre"`
	publication_date string `json:"publication_date"`
	created_at time.Time `json:"created_at"`
	updated_at time.Time `json:"updated_at"`
}

func main() {
	router := gin.Default()
	port := os.Getenv("PORT")
	err := godotenv.Load()
	
	if port == "" {
		port = "8080"
	}
	
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message": "Hello world!",
		})
	})

	router.Run(":" + port)
}