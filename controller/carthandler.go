package controller

import (
	"net/http"
	"text/template"

	"github.com/shuwenhe/shuwen-bookstore/utils"

	"github.com/shuwenhe/shuwen-bookstore/model"

	"github.com/shuwenhe/shuwen-bookstore/dao"
)

// AddBook2Cart Add book to cart
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	flag, session := dao.IsLogin(r) // Determine whether to log in
	if flag {                       // Already logged in
		bookID := r.FormValue("bookId") // Get bookId
		book, _ := dao.GetBookByID(bookID)
		userID := session.UserID
		cart, _ := dao.GetCartByUserID(userID) // Determine whether there is a current cart in the database
		if cart != nil {                       // Already have a cart
			carItem, _ := dao.GetCartItemByBookIDAndCartID(bookID, cart.CartID) // The current user already has a cart and needs to determine whether the book already exists in the current cart
			if carItem != nil {                                                 // The book is already in the cartItem of the cart, just add 1 to the data of the cartItem in the cart
				cts := cart.CartItems   // 1.Get all the cartItems in the cart slice
				for _, v := range cts { // 2.For range get every cartItems
					if v.Book.ID == carItem.Book.ID { // 3.Find the current cartItem
						v.Count = v.Count + 1  // Add 1 to the number of books in the cart
						dao.UpdateBookCount(v) // Update the number of books of the cartItem in the database
					}
				}
			} else { // The cartItem in the cart does not have the book. At this time, you need to create a cartItem and add it to the cart
				cartItem := &model.CartItem{
					Book:   book,
					Count:  1,
					CartID: cart.CartID,
				}
				cart.CartItems = append(cart.CartItems, cartItem) // 将购物项添加到当前cart的slice中
				dao.AddCartItem(cartItem)                         // 将新创建的购物项添加到数据中
			}
			dao.UpdateCart(cart) // 不管之前购物车中是否有当前对应的购物项，都需要更新购物车中的图书的总数量和总价格
		} else { // 当前用户没有购物车，需要创建一个购物车并添加到数据库中
			cartID := utils.CreateUUID() // 生成购物车id
			cart := &model.Cart{         // 1.创建购物车，可以不一次赋值
				CartID: cartID,
				UserID: userID,
			}
			var cartItems []*model.CartItem // 2.创建购物车中的购物项,声明一个cartItems类型的slice
			cartItem := &model.CartItem{    // 购物项CartItemID是自增的不用设置
				Book:   book,
				Count:  1, // count有了amount也就有了
				CartID: cartID,
			}
			cartItems = append(cartItems, cartItem) // 将cartItem放入slice
			cart.CartItems = cartItems              // 3.将cartItems放入slice中
			dao.AddCart(cart)                       // 4.将cart保存到database
		}
		w.Write([]byte("You just added" + book.Title + "to cart"))
	} else { // Not logged in
		w.Write([]byte("please login first!"))
	}
}

// GetCartInfo Get cart information based on user ID
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r) // Determine whether you have logged in
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID) // Get the corresponding cart from the database according to the user's id
	if cart != nil {                       // Determine whether the shopping cart is empty
		cart.UserName = session.UserName
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html")) // Parse the template file
		t.Execute(w, cart)                                                    // execute
	} else { // The user does not have a cart
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html")) // Parse the template file
		t.Execute(w, session)                                                 // Execute template
	}
}

// DeleteCart delete cart
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	cartID := r.FormValue("cartId") // Get the ID of the cart to be deleted
	dao.DeleteCartByCartID(cartID)  // empty cart
	GetCartInfo(w, r)               // Call GetCartInfo function to query information again
}
