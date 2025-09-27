package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	database := Database{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		Hostname: os.Getenv("HOSTNAME"),
		Port:     StringToInteger(os.Getenv("PORT")),
		Database: os.Getenv("DATABASE"),
	}
	if err := database.init(); err != nil {
		panic(err)
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	signRoute(router)
	router.Run(":8080")
}
