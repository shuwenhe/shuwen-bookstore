package controller

import (
	"net/http"
	"text/template"
	"time"

	"github.com/shuwenhe/shuwen-bookstore/utils"

	"github.com/shuwenhe/shuwen-bookstore/model"

	"github.com/shuwenhe/shuwen-bookstore/dao"
)

// Checkout checkout
func Checkout(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r) // Get the session
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	orderID := utils.CreateUUID() // Generate order number
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	order := &model.Order{ // Create order
		OrderID:     orderID,
		CreateTime:  timeStr,
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserID:      int64(userID),
	}
	dao.AddOrder(order)           // Save the order to the database
	cartItems := cart.CartItems   // Get the cartItem in the cart
	for _, v := range cartItems { // for range to get every cartItem
		orderItem := &model.OrderItem{ // create orderItem
			Count:   v.Count,
			Amount:  v.Amount,
			Title:   v.Book.Title,
			Author:  v.Book.Author,
			Price:   v.Book.Price,
			ImgPath: v.Book.ImgPath,
			OrderID: orderID,
		}
		dao.AddOrderItem(orderItem)
		book := v.Book // Update the stocks and sales of books in the current cartItems
		book.Sales = book.Sales + v.Count
		book.Stock = book.Stock - v.Count
		dao.UpdateBook(book) // Update the information of the books
	}
	dao.DeleteCartByCartID(cart.CartID)                                       // Delete the cart
	session.Order = order                                                     // Set the order number to the session
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html")) // Parse the template
	t.Execute(w, session)                                                     // Execute
}

// GetOrders Get all orders
func GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, _ := dao.GetOrders()
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	t.Execute(w, orders)
}

func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	orderID := r.FormValue("orderId")
	orderItems, _ := dao.GetOrderItemsByOrderID(orderID)
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	t.Execute(w, orderItems)
}

func GetOrderByUserID(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	userID := session.UserID
	orders, _ := dao.GetOrderByUserID(userID)
	session.Orders = orders
	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	t.Execute(w, session)
}

func SendOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.FormValue("orderId")
	dao.UpdateOrderState(orderID, 1)
	GetOrders(w, r)
}

// TakeOrder take order
func TakeOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.FormValue("orderId")
	dao.UpdateOrderState(orderID, 2)
	GetOrderByUserID(w, r)
}
