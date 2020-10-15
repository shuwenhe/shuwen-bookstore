package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/shuwenhe/shuwen-bookstore/model"

	"github.com/shuwenhe/shuwen-bookstore/dao"
)

// DeleteCartItem delete cartItems
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemId")
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	_, session := dao.IsLogin(r)           // Get session
	userID := session.UserID               // Get user's ID
	cart, _ := dao.GetCartByUserID(userID) // Get the cart base on the userID
	cartItems := cart.CartItems            // Get the cartItems in the cart
	for k, v := range cartItems {          // Range to get every cartItem
		if v.CartItemID == iCartItemID { // Find the cart Item to delete
			cartItems = append(cartItems[:k], cartItems[k+1:]...) // Remove the current cartItem from the cartItems's slice
			cart.CartItems = cartItems                            // Assign the slice after deleting the item to the slice in the cart again
			dao.DeleteCartItemByID(cartItemID)                    // Delete the current cartItems from the database
		}
	}
	dao.UpdateCart(cart) // Update total count and total amount of books in the cart
	GetCartInfo(w, r)
}

// UpdateCartItem update cartItem
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemId")
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	bookCount := r.FormValue("bookCount")
	iBookCount, _ := strconv.ParseInt(bookCount, 10, 64)
	_, session := dao.IsLogin(r)           // Get the session
	userID := session.UserID               // Get user's ID
	cart, _ := dao.GetCartByUserID(userID) // Get the cart of the user
	cartItems := cart.CartItems            // Get all cartItmes in the cart
	for _, v := range cartItems {          // Range the cartItems slice get every cartItem
		if v.CartItemID == iCartItemID { // Find the update cartItem
			v.Count = iBookCount
			dao.UpdateBookCount(v)
		}
	}
	dao.UpdateCart(cart)
	// GetCartInfo(w, r)
	cart, _ = dao.GetCartByUserID(userID)
	totalCount := cart.TotalCount
	totalAmount := cart.TotalAmount
	var amount float64
	cIs := cart.CartItems
	for _, v := range cIs {
		if iCartItemID == v.CartItemID {
			amount = v.Amount
		}
	}
	data := model.Data{
		Amount:      amount,
		TotalAmount: totalAmount,
		TotalCount:  totalCount,
	}
	json, _ := json.Marshal(data)
	w.Write(json)
}
