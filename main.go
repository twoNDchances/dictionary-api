package main

import (
	"dictionary-api/models"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	database := models.Database{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		Hostname: os.Getenv("HOSTNAME"),
		Port:     StringToInteger(os.Getenv("PORT")),
		Database: os.Getenv("DATABASE"),
	}
	if err := database.Init(); err != nil {
		panic(err)
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	signRoute(router)
	router.Run(":8080")
}
