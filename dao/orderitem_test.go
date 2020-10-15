package dao

import (
	"fmt"
	"testing"
)

func testGetOrderItemsByOrderID(t *testing.T) {
	orderItems, _ := GetOrderItemsByOrderID("942e5255-edad-4ffe-4921-ed468c84069d")
	for _, orderItem := range orderItems {
		fmt.Println("orderItem = ", orderItem)
	}
}
