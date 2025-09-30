package main

import (
	"dictionary-api/controllers"

	"github.com/gin-gonic/gin"
)

func signRoute(engine *gin.Engine) {
	apiGroup := engine.Group("/api")
	{
		v1Group := apiGroup.Group("/v1")
		{
			dictionaryGroup := v1Group.Group("/dictionaries")
			{
				dictionaryGroup.GET("", controllers.List)
				dictionaryGroup.GET("/:id", controllers.Show)
				dictionaryGroup.POST("", controllers.Create)
				dictionaryGroup.PATCH("/:id", controllers.Update)
				dictionaryGroup.DELETE("/:id", controllers.Delete)
			}
		}
	}
}
