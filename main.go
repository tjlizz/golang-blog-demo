package main

import (
	"golang-blog/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
}
func main() {

	config.ConnectMysql()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
