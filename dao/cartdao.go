package dao

import (
	"github.com/shuwenhe/shuwen-bookstore/utils"

	"github.com/shuwenhe/shuwen-bookstore/model"
)

// AddCart 向购物车表中插入购物车
func AddCart(cart *model.Cart) error {
	sql := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"                     // id为什么需要？                    // sql
	_, err := utils.Db.Exec(sql, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID) // 执行
	if err != nil {
		return err
	}
	cartItems := cart.CartItems          // 获取购物车中所有购物项
	for _, cartItem := range cartItems { // 遍历得到每一个购物项
		AddCartItem(cartItem) // 将购物项保存到数据库中
	}
	return nil
}

// GetCartByUserID 通过用户ID去数据库中查询购物车
func GetCartByUserID(userID int) (*model.Cart, error) {
	sql := "select id,total_count,total_amount,user_id from carts where user_id = ?" // sql
	row := utils.Db.QueryRow(sql, userID)                                            // 执行sql
	cart := &model.Cart{}
	err := row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	if err != nil {
		return nil, err
	}
	cartItems, _ := GetCartItemsByCartID(cart.CartID) // 通过cartID获取cartItems，对数据的整合，全局观，需要哪个func，然后再写，架构思维(全局观(预判(先见之明(准备好了再干))))要有，首先要明确做什么，然后才能谈怎么做
	cart.CartItems = cartItems                        // 将获取的cartItems的slice存入cart结构体中的cartItems
	return cart, nil                                  //返回
}

// UpdateCart Update the total number and total amount of books in the shopping cart
func UpdateCart(cart *model.Cart) error {
	sql := "update carts set total_count = ?,total_amount = ? where id = ?"                // sql
	_, err := utils.Db.Exec(sql, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID) // execute
	if err != nil {
		return err
	}
	return nil
}

// DeleteCartByCartID Delete cart base on the ID of the cart
func DeleteCartByCartID(cartID string) error {
	err := DeleteCartItemByCartID(cartID) // Need to delete all cartItems before delete cart
	if err != nil {
		return err
	}
	sql := "delete from carts where id = ?" // write sql statement
	_, err2 := utils.Db.Exec(sql, cartID)   // execute
	if err2 != nil {
		return err2
	}
	return nil
}
