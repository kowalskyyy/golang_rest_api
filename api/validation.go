package api

import "errors"

func validateOrder(order order) error {
	if order.CustomerID == "" {
		return errors.New("customer ID is required")
	}
	if order.OrderID == "" {
		return errors.New("order ID is required")
	}
	if len(order.Items) == 0 {
		return errors.New("at least one item is required")
	}
	for _, item := range order.Items {
		if item.ItemID == "" {
			return errors.New("incorrect itemID")
		}
	}
	// Add more validation rules as needed
	return nil
}
