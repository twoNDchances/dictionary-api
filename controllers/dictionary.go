package controllers

import (
	"dictionary-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	dictionaryList, status, err := models.List()
	if err != nil {
		c.AbortWithStatusJSON(status, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"data":    dictionaryList,
		"message": "success",
	})
}

func Create(c *gin.Context) {
	var dictionary models.Dictionary
	if err := c.ShouldBindJSON(&dictionary); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	dictionaryCreate, status, err := dictionary.Create()
	if err != nil {
		c.AbortWithStatusJSON(status, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"data":    dictionaryCreate,
		"message": "success",
	})
}

func Show(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": "invalid id",
		})
		return
	}
	dictionaryShow, status, err := models.Show(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(status, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"data":    dictionaryShow,
		"message": "success",
	})
}

func Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": "invalid id",
		})
		return
	}

	var dictionary models.Dictionary
	if err := c.ShouldBindJSON(&dictionary); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	dictionaryUpdate, status, err := dictionary.Update(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(status, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"data":    dictionaryUpdate,
		"message": "success",
	})
}

func Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"data":    nil,
			"message": "invalid id",
		})
		return
	}
	status, err := models.Delete(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(status, gin.H{
			"data":    nil,
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"data":    nil,
		"message": "success",
	})
}
