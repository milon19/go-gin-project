package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"healthousedemo/cmd/app/config"
	"healthousedemo/internal/app/adapter"
	authConfig "healthousedemo/internal/app/adapter/config"
	db "healthousedemo/internal/app/adapter/db/connections"
	"log"
	"net/http"
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
	authConfig.LoadConfig()
	db.Connect()
}

func main() {
	r := gin.Default()
	r = adapter.Routes(r)
	err := r.SetTrustedProxies([]string{})
	if err != nil {
		return
	}
	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"payload": "Method Not Allowed"})
	})
	if runError := r.Run(":" + config.GinPort); runError != nil {
		log.Println(runError)
	}
}
