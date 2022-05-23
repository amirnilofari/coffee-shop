package main

import (
	"fmt"

	"github.com/coffee-shop/models"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Start Coffe Shop project")

	server := gin.Default()

	//connect to database
	models.ConnectDB()

	v1 := server.Group("/v1")
	{
		v1.GET("/order", OrderDrink)
		v1.GET("/menu", GetMenu)
	}

	server.Run()

}
