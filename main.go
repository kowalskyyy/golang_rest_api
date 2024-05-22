package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

type items struct {
	ItemID  string  `json:"itemId"`
	CostEur float32 `json:"costEur"`
}

type order struct {
	CustomerID string  `json:"customerId"`
	OrderID    string  `json:"orderId"`
	Timestamp  string  `json:"timestamp"`
	Items      []items `json:"items"`
}

var orders = []order{
	{CustomerID: "01",
		OrderID:   "50",
		Timestamp: "1637245070513",
		Items: []items{{
			ItemID: "20201", CostEur: 2},
		},
	},
	{CustomerID: "01",
		OrderID:   "50",
		Timestamp: "1637245070513",
		Items: []items{{
			ItemID: "20201", CostEur: 2},
		},
	}}

func getOrders(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, orders)
}

func main() {
	fmt.Printf("elo mordo")
	router := gin.Default()
	router.GET("/get-orders", getOrders)
	router.Run("localhost:8080")
}
