package model

// Cart 购物车结构体
type Cart struct {
	CartID      string      // 购物车ID
	CartItems   []*CartItem // 购物车中的购物项
	TotalCount  int64       // 购物车中图书的总数量，通过计算得到
	TotalAmount float64     // 购物车中图书总金额，通过计算得到
	UserID      int         // 当前购物车所属用户
	UserName    string
}

// GetTotalCount 获取图书的总数量
func (cart *Cart) GetTotalCount() int64 {
	var totalCount int64
	for _, v := range cart.CartItems { // 遍历购物车中的购物项目slice
		totalCount = totalCount + v.Count
	}
	return totalCount
}

// GetTotalAmount 获取图书总金额
func (cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	for _, v := range cart.CartItems { // 遍历购物车中的购物项目slice
		totalAmount = totalAmount + v.GetAmount()
	}
	return totalAmount
}
