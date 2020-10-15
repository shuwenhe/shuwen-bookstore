package model

// Session struct
type Session struct { // Go语言中struct是值类型，用*效率高
	SessionID string
	UserName  string
	UserID    int // 外键，关联User结构体
	Cart      *Cart
	Order     *Order
	Orders    []*Order
}
