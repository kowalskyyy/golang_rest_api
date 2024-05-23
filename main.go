package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

type item struct {
	ItemID  string  `json:"itemId"`
	CostEur float32 `json:"costEur"`
}

type customerItem struct {
	CustomerID string `json:"customerId"`
	item
}

type order struct {
	CustomerID string `json:"customerId"`
	OrderID    string `json:"orderId"`
	Timestamp  string `json:"timestamp"`
	Items      []item `json:"items"`
}

type customerSummary struct {
	CustomerID          string  `json:"customerId"`
	NbrOfPurchasedItems int16   `json:"nbrOfPurchasedItems"`
	TotalAmountEur      float32 `json:"totalAmountEur"`
}

var orders = []order{}

func getOrders(context *gin.Context) {
	if len(orders) == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "There are no orders in the system. Please submit orders first."})
		return
	}
	context.IndentedJSON(http.StatusOK, orders)
}

func submitOrders(context *gin.Context) {
	var newOrders []order
	if err := context.BindJSON(&newOrders); err != nil {

		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Your order data is incorrect"})
		return
	}
	orders = append(orders, newOrders...)
	context.IndentedJSON(http.StatusOK, newOrders)
}

func getCustomersSummary(context *gin.Context) {
	summaryMap := make(map[string]*customerSummary)
	var summaries []customerSummary
	for _, order := range orders {
		if _, exists := summaryMap[order.CustomerID]; !exists {
			summaryMap[order.CustomerID] = &customerSummary{
				CustomerID:          order.CustomerID,
				NbrOfPurchasedItems: 0,
				TotalAmountEur:      0,
			}
		}

		summary := summaryMap[order.CustomerID]
		for _, item := range order.Items {
			summary.TotalAmountEur += item.CostEur
		}
		summary.NbrOfPurchasedItems++
	}
	for _, summary := range summaryMap {
		summaries = append(summaries, *summary)
	}

	if len(summaries) != 0 {
		context.IndentedJSON(http.StatusOK, summaries)
		return
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No summaries available. There are no orders in the system."})
}

func getItemsByCustomer(context *gin.Context) {
	id := context.Param("id")
	var result []customerItem
	if len(orders) != 0 {
		for _, order := range orders {
			if order.CustomerID == id {
				for _, item_ := range order.Items {
					extended := customerItem{
						CustomerID: id,
						item: item{
							ItemID:  item_.ItemID,
							CostEur: item_.CostEur,
						},
					}
					result = append(result, extended)
				}
			}
		}
		if len(result) != 0 {
			context.IndentedJSON(http.StatusOK, gin.H{"items": result})
			return
		} else {
			context.IndentedJSON(http.StatusNotFound, gin.H{"message": "There are no orders associated with this customer ID"})
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "There are no orders in the system. Please submit orders first."})
}

func main() {
	fmt.Printf("elo mordo")
	router := gin.Default()
	router.GET("/get-orders", getOrders)
	router.GET("/get-items/:id", getItemsByCustomer)
	router.GET("/get-summary", getCustomersSummary)
	router.POST("/submit-orders", submitOrders)
	router.Run("localhost:8080")
}
