package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/get-orders", getOrders)
	router.GET("/get-items/:id", getItemsByCustomer)
	router.GET("/get-summary", getCustomersSummary)
	router.POST("/submit-orders/:strict", submitOrders)
	router.POST("/submit-orders", submitOrders)

}
