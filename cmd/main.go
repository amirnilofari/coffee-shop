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

	server.GET("/menu", GetMenu)
	server.Run()

}
