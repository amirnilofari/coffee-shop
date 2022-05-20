package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/coffee-shop/models"
)

func main() {
	fmt.Println("Start Coffe Shop project")

	collection := []models.Drink{
		{ID: "1", Name: "Cappacino", Type: "warm", Price: 3.99},
		{ID: "2", Name: "Icepack", Type: "cold", Price: 5.99},
		{ID: "3", Name: "Tea", Type: "warm", Price: 2.19},
		{ID: "4", Name: "Ice tea", Type: "cold", Price: 2.29},
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter order number:")
	//var orderNum int

	for _, v := range collection {
		fmt.Println(v.ID, v.Name)

	}

	fmt.Print("-> ")
	order, _ := reader.ReadString('\n')
	// convert CRLF to LF
	order = strings.Replace(order, "\n", "", -1)
	var price1 float64
	for _, v := range collection {
		if order == v.ID {
			price1 = v.Price
			fmt.Println("your order price is : ", price1)
		}

	}

	fmt.Println("another order ? Yes/No")
	fmt.Print("-> ")
	anotherOrder, _ := reader.ReadString('\n')
	// convert CRLF to LF
	anotherOrder = strings.Replace(anotherOrder, "\n", "", -1)
	if anotherOrder == "Yes" {
		for _, v := range collection {
			fmt.Println(v.ID, v.Name)

		}
		fmt.Print("-> ")
		order2, _ := reader.ReadString('\n')
		// convert CRLF to LF
		order2 = strings.Replace(order2, "\n", "", -1)
		for _, v := range collection {
			if order2 == v.ID {
				fmt.Println("your order price is : ", v.Price+price1)
			}

		}
	}

	// for _, v := range collection {
	// 	if v.ID == 1 {
	// 		fmt.Println("price your order is", v.Price)
	// 	}
	// 	if v.ID == 2 {
	// 		fmt.Println("price your order is", v.Price)
	// 	}
	// 	if v.ID == 3 {
	// 		fmt.Println("price your order is", v.Price)
	// 	}
	// 	if v.ID == 4 {
	// 		fmt.Println("price your order is", v.Price)
	// 	}
	// }
	//	fmt.Println(collection)

}
