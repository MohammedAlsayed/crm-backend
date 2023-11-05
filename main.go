package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(corsMiddleware())

	r.POST("/api/login", func(c *gin.Context) {
		// read request body
		username := c.PostForm("username")
		password := c.PostForm("password")

		// print username and password
		fmt.Printf("username: %s, passowrd: %s", username, password)

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	r.Run()
}
