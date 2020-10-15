package dao

import (
	"fmt"
	"testing"

	"github.com/shuwenhe/shuwen-bookstore/model"
)

func testAddCart(t *testing.T) {
	book := &model.Book{ // 设置购买的第1本图书
		ID:    1,
		Price: 27.20,
	}

	book2 := &model.Book{ // 设置购买的第2本图书
		ID:    2,
		Price: 23.00,
	}

	var cartItems []*model.CartItem // create cartItem slice
	cartItem := &model.CartItem{    // 创建第1个购物项
		Book:   book,
		Count:  10,
		CartID: "123456",
	}
	cartItems = append(cartItems, cartItem) // cartItem append to cartItems slice
	cartItem2 := &model.CartItem{           // 创建第2个购物项
		Book:   book2,
		Count:  10,
		CartID: "123456",
	}
	cartItems = append(cartItems, cartItem2) // cartItem append to cartItems slice
	cart := &model.Cart{                     // create cart
		CartID:    "123456",
		CartItems: cartItems,
		UserID:    1,
	}
	AddCart(cart)
}

func testGetCartByUserID(t *testing.T) {
	cart, _ := GetCartByUserID(3)
	fmt.Println("cart = ", cart)
}

func testDeleteCartByCartID(t *testing.T) {
	DeleteCartByCartID("4f702f4e-eff8-40ce-5886-dc443b816a94")
}
