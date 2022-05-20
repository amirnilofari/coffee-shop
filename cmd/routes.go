package main

import (
	"net/http"

	"github.com/coffee-shop/models"
	"github.com/gin-gonic/gin"
)

func GetMenu(c *gin.Context) {
	collection := []models.Drink{
		{ID: "1", Name: "Cappacino", Type: "warm", Price: 3.99},
		{ID: "2", Name: "Icepack", Type: "cold", Price: 5.99},
		{ID: "3", Name: "Tea", Type: "warm", Price: 2.19},
		{ID: "4", Name: "Ice tea", Type: "cold", Price: 2.29},
	}

	c.JSON(http.StatusOK, collection)

}
