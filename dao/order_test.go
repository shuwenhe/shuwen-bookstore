package dao

import (
	"fmt"
	"testing"
	"time"

	"github.com/shuwenhe/shuwen-bookstore/model"
)

func testAddOrder(t *testing.T) {
	orderID := "15010729356"
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	order := &model.Order{ // create order
		OrderID:     orderID,
		CreateTime:  timeStr,
		TotalCount:  2,
		TotalAmount: 400,
		State:       0,
		UserID:      1,
	}

	orderItem := &model.OrderItem{ // create orderItems
		Count:   1,
		Amount:  300,
		Title:   "Unix",
		Author:  "Ken",
		Price:   300,
		ImgPath: "static/img/default.jpg",
		OrderID: orderID,
	}

	orderItem2 := &model.OrderItem{
		Count:   1,
		Amount:  100,
		Title:   "Linux",
		Author:  "Linus",
		Price:   100,
		ImgPath: "static/img/default.jpg",
		OrderID: orderID,
	}
	AddOrder(order) // Save order

	AddOrderItem(orderItem) // Save OrderItem
	AddOrderItem(orderItem2)
}

func testGetOrders(t *testing.T) {
	orders, _ := GetOrders()
	for _, order := range orders {
		fmt.Println("order = ", order)
	}
}

func testGetOrderByUserID(t *testing.T) {
	orders, _ := GetOrderByUserID(1)
	fmt.Println("orders=", orders)
	for _, order := range orders {
		fmt.Println("order=", order)
	}
}

func TestUpdateOrderState(t *testing.T) {
	UpdateOrderState("942e5255-edad-4ffe-4921-ed468c84069d", 1)
}
