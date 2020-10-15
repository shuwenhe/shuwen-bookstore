package dao

import (
	"fmt"
	"testing"

	"github.com/shuwenhe/shuwen-bookstore/model"
)

func testGetCartItemByBookID(t *testing.T) {
	cartItem, _ := GetCartItemByBookIDAndCartID("1", "123456")
	fmt.Println("图书id = 1的cartItem是", cartItem)
}

func testGetCartItemByCartID(t *testing.T) {
	cartItems, _ := GetCartItemsByCartID("123456")
	for k, cartItem := range cartItems {
		fmt.Printf("第%v个购物项是:%v", k+1, cartItem)
		fmt.Println()
	}
}

func testUpdateBookCount(t *testing.T) {
	cartItem := &model.CartItem{}
	UpdateBookCount(cartItem) // Update the number of books in the cartItem according to the id of the book, the id of the cart and the number of books
}

func TestDeleteCartItemByCartID(t *testing.T) {
	DeleteCartItemByCartID("4f702f4e-eff8-40ce-5886-dc443b816a94")
}

func TestDeleteCartItemByID(t *testing.T) {
	DeleteCartItemByID("40")
}
