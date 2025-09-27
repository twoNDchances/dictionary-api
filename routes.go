package main

import "github.com/gin-gonic/gin"

func signRoute(engine *gin.Engine)  {
	apiGroup := engine.Group("/api")
	{
		v1Group := apiGroup.Group("/v1")
		{
			dictionaryGroup := v1Group.Group("/dictionaries")
			{
				dictionaryGroup.GET("", list)
				dictionaryGroup.GET("/:key", show)
				dictionaryGroup.POST("", create)
				dictionaryGroup.PATCH("", update)
				dictionaryGroup.DELETE("", delete)
			}
		}
	}
}
