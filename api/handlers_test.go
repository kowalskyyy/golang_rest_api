package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetOrders(t *testing.T) {
	router := gin.Default()
	router.GET("/get-orders", getOrders)
	req, err := http.NewRequest("GET", "/get-orders", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
	var response map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Error parsing JSON response: %v", err)
	}
	expectedMessage := "There are no orders in the system. Please submit orders first."
	if message, ok := response["message"]; !ok || message != expectedMessage {
		t.Errorf("Handler returned unexpected message: got %q want %q", message, expectedMessage)
	}
}

func TestSubmitOrdersValid(t *testing.T) {
	router := gin.Default()
	router.POST("/submit-orders", submitOrders)
	validOrder := order{
		CustomerID: "123",
		OrderID:    "456",
		Timestamp:  "2024-05-23T12:00:00Z",
		Items: []item{
			{ItemID: "item1", CostEur: 10.5},
			{ItemID: "item2", CostEur: 20.3},
		},
	}
	jsonValidOrder, err := json.Marshal([]order{validOrder})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/submit-orders?strict=false", bytes.NewBuffer(jsonValidOrder))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestSubmitOrdersInvalid(t *testing.T) {
	router := gin.Default()
	router.POST("/submit-orders", submitOrders)
	invalidOrder := order{
		CustomerID: "123",
		// Missing OrderID, Timestamp, and Items
	}
	jsonInvalidOrder, err := json.Marshal([]order{invalidOrder})
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/submit-orders", bytes.NewBuffer(jsonInvalidOrder))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusPartialContent && status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusPartialContent)
	}
	expected := "Some of the orders had a faulty format."
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Handler returned unexpected body: got %q want response containing %q", rr.Body.String(), expected)
	}
}

func TestGetOrdersWithOrderrs(t *testing.T) {
	router := gin.Default()
	router.GET("/get-orders", getOrders)
	req, err := http.NewRequest("GET", "/get-orders", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}
