package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"healthousedemo/cmd/app/config"
	db "healthousedemo/internal/app/adapter/db/connections"
	"log"
)

func init() {
	initEnv()
}

func initEnv() {
	log.Printf("Loading environment setting")
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No local env file. Using global OS environment variables")
	}
	config.SetEnvironment()
	db.Connect()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	if runError := r.Run(":" + config.GinPort); runError != nil {
		log.Println(runError)
	}
}
