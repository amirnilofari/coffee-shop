package main

import (
	"net/http"

	"github.com/coffee-shop/models"
	"github.com/gin-gonic/gin"
)

func GetMenu(c *gin.Context) {
	var drinks []models.Drink

	models.DB.Find(&drinks)

	c.JSON(http.StatusOK, gin.H{"data": drinks})

}
