package main

import "github.com/gin-gonic/gin"

func list(context *gin.Context) {
	dictionary := Dictionary{}
	dictionaryList, status, err := dictionary.list()
	if err != nil {
		context.AbortWithStatusJSON(
			status,
			gin.H{
				"data": nil,
				"message": err.Error(),
			},
		)
		return
	}
	context.JSON(
		status,
		gin.H{
			"data": dictionaryList,
			"message": "success",
		},
	)
}

func create(context *gin.Context) {
	var dictionary Dictionary
	if err := context.ShouldBindBodyWithJSON(&dictionary); err != nil {
		context.AbortWithStatusJSON(
			400,
			gin.H{
				"data": nil,
				"message": err.Error(),
			},
		)
		return
	}
	dictionaryCreate, status, err := dictionary.create()
	if err != nil {
		context.AbortWithStatusJSON(
			status,
			gin.H{
				"data": nil,
				"message": err.Error(),
			},
		)
		return
	}
	context.JSON(
		status,
		gin.H{
			"data": dictionaryCreate,
			"message": "success",
		},
	)
}

func show(context *gin.Context) {
	var dictionary Dictionary
	if err := context.ShouldBindUri(&dictionary); err != nil {
		context.AbortWithStatusJSON(
			400,
			gin.H{
				"data": nil,
				"message": err.Error(),
			},
		)
		return
	}
	dictionaryShow, status, err := dictionary.show()
	if err != nil {
		context.AbortWithStatusJSON(
			status,
			gin.H{
				"data": nil,
				"message": err.Error(),
			},
		)
		return
	}
	context.JSON(
		status,
		gin.H{
			"data": dictionaryShow,
			"message": "success",
		},
	)
}

func update(context *gin.Context) {
	var dictionary Dictionary
	if err := context.ShouldBindBodyWithJSON(&dictionary); err != nil {
		context.AbortWithStatusJSON(
			400,
			gin.H{
				"data": nil,
				"message": err.Error(),
			},
		)
		return
	}
	dictionaryUpdate, status, err := dictionary.update()
	if err != nil {
		context.AbortWithStatusJSON(
			status,
			gin.H{
				"data": nil,
				"message": err.Error(),
			},
		)
		return
	}
	context.JSON(
		status,
		gin.H{
			"data": dictionaryUpdate,
			"message": "success",
		},
	)
}

func delete(context *gin.Context) {
	var dictionary Dictionary
	if err := context.ShouldBindBodyWithJSON(&dictionary); err != nil {
		context.AbortWithStatusJSON(
			400,
			gin.H{
				"data": nil,
				"message": err.Error(),
			},
		)
		return
	}
	dictionaryDelete, status, err := dictionary.delete()
	if err != nil {
		context.AbortWithStatusJSON(
			status,
			gin.H{
				"data": nil,
				"message": err.Error(),
			},
		)
		return
	}
	context.JSON(
		status,
		gin.H{
			"data": dictionaryDelete,
			"message": "success",
		},
	)
}
