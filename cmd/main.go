package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Start Coffe Shop project")

	server := gin.Default()
	server.GET("/menu", GetMenu)
	server.Run()

}
