package main

import (
	"goosint/controllers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	router := gin.Default()
	router.POST("/phone", controllers.PhoneController)
	router.POST("/email", controllers.EmailController)
	router.POST("/username", controllers.UsernameController)
	router.POST("/facebook", controllers.FacebookController)
	router.POST("/ip", controllers.IPController)
	router.Run(":8088")
}
