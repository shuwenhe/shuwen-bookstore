package dao

import (
	"github.com/shuwenhe/shuwen-bookstore/utils"

	"github.com/shuwenhe/shuwen-bookstore/model"
)

// AddCartItem 向购物项表中插入购物项
func AddCartItem(cartItem *model.CartItem) error {
	sql := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"                         // 为什么不插入id                       // id为什么不需要                        // sql
	_, err := utils.Db.Exec(sql, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID) // 执行
	if err != nil {
		return err
	}
	return nil
}

// GetCartItemByBookIDAndCartID 根据图书的ID获取对应的购物项
func GetCartItemByBookIDAndCartID(bookID string, cartID string) (*model.CartItem, error) {
	sql := "select id,count,amount,cart_id from cart_items where book_id = ? and cart_id = ?" // sql
	row := utils.Db.QueryRow(sql, bookID, cartID)                                             // 执行
	cartItem := &model.CartItem{}                                                             // 创建cartItem
	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil, err
	}
	book, err := GetBookByID(bookID) // Query book information according to the id of the book
	if err != nil {
		return nil, err
	}
	cartItem.Book = book // Set the book to the cartItem
	return cartItem, nil
}

// GetCartItemsByCartID 根据购物车的ID获取对应的所有的购物项
func GetCartItemsByCartID(cartID string) ([]*model.CartItem, error) {
	sql := "select id,count,amount,book_id,cart_id from cart_items where cart_id = ?" // sql
	rows, err := utils.Db.Query(sql, cartID)                                          // execute
	if err != nil {
		return nil, err
	}
	var cartItems []*model.CartItem
	for rows.Next() {
		var bookID string             // 设置一个变量接收bookId
		cartItem := &model.CartItem{} // 创建cartItem
		err2 := rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &bookID, &cartItem.CartID)
		if err2 != nil {
			return nil, err2
		}
		book, _ := GetBookByID(bookID) // 根据bookID获取图书信息
		cartItem.Book = book           // 将book设置到cartItems中
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

// UpdateBookCount Update the number of books in the cartItem according to the id of the book, the id of the cart and the number of books
func UpdateBookCount(cartItem *model.CartItem) error {
	sql := "update cart_items set count = ?,amount = ? where book_id = ? and cart_id = ?"                 // sql
	_, err := utils.Db.Exec(sql, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID) // execute
	if err != nil {
		return err
	}
	return nil
}

// DeleteCartItemByCartID Delete CartItem base on the ID of the cart
func DeleteCartItemByCartID(cartID string) error {
	sql := "delete from cart_items where cart_id = ?" // sql
	_, err := utils.Db.Exec(sql, cartID)              // execute
	if err != nil {
		return err
	}
	return nil
}

// DeleteCartItemByID delete cart items base on their ID
func DeleteCartItemByID(cartItemID string) error {
	sql := "delete from cart_items where id = ?"
	_, err := utils.Db.Exec(sql, cartItemID)
	if err != nil {
		return err
	}
	return nil
}
