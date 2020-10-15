package model

// OrderItem struct
type OrderItem struct {
	OrderItemID int64
	Count       int64
	Amount      float64
	Title       string
	Author      string
	Price       float64
	ImgPath     string
	OrderID     string
}
