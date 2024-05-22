package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

{
	"customerId": "01",
	"orderId": "50",
	"timestamp": "1637245070513",
	"items": [
	{
	
	"itemId": "20201",
	"costEur": 2
	}
	]
	}

type items struct {
	ItemID		string	`json:"itemId"`
	CostEur		float32	`json:"costEur"`
}

type order struct {
	CustomerID	string	`json:"customerId"`
	OrderID		string	`json:"orderId"`
	Timestamp	string	`json:"timestamp"`
	Items		items	`json:"items"`

}

func main() {
	fmt.Printf("elo mordo")
	router := gin.Default()
	router.Run("localhost:8080")
}
