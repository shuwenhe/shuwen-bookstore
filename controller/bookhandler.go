package controller

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/shuwenhe/shuwen-bookstore/model"

	"github.com/shuwenhe/shuwen-bookstore/dao"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, _ := dao.GetBooks()
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, books)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")
	dao.DeleteBook(bookId)
	GetPageBooks(w, r)
}

// ToUpdateBookPage 去更新或添加图书的页面
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookId")    // 获取要更新的图书的id
	book, _ := dao.GetBookByID(bookID) // 调取bookdao中获取图书的函数
	if book.ID > 0 {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html")) // 解析模板
		t.Execute(w, book)                                                            // 执行
	} else {
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html")) // 解析模板
		t.Execute(w, "")                                                              // 执行
	}
}

// UpdateOrAddBook 更新或添加图书
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	bookID := r.PostFormValue("bookId") // 获取图书信息
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	fPrice, _ := strconv.ParseFloat(price, 64) // 将价格、销量、库存进行转换
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)
	iBookID, _ := strconv.ParseInt(bookID, 10, 0)
	book := &model.Book{ // 创建book
		ID:      int(iBookID),
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   iSales,
		Stock:   iStock,
		ImgPath: "/static/img/default.jpg",
	}
	if book.ID > 0 {
		dao.UpdateBook(book) // 调用更新图书的函数
	} else {
		dao.AddBook(book)
	}
	GetPageBooks(w, r)
}

// GetPageBooks 获取分页图书
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPageBooks(pageNo)
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, page)
}

// GetPageBooksByPrice 获取带分页和价格范围的图书
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")
	minPrice := r.FormValue("min") // FormValue与PostFormValue的区别？
	maxPrice := r.FormValue("max") // 超链接中不是POST请求
	if pageNo == "" {
		pageNo = "1"
	}
	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		page, _ = dao.GetPageBooks(pageNo) // 调用bookdao中获取带分页的图书的函数
	} else {
		page, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
		page.MinPrice = minPrice // 将价格范围设置到page中
		page.MaxPrice = maxPrice
	}
	flag, session := dao.IsLogin(r) // 调用IsLogin判断用户是否已经登录
	if flag {
		page.IsLogin = true // 用户已经登录，设置page中的IsLogin字段和Username的字段值
		page.Username = session.UserName
	}
	t := template.Must(template.ParseFiles("views/index.html")) // 解析模板
	t.Execute(w, page)                                          // 执行
}

// Indexhandler 获取首页
func Indexhandler(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo") // 获取页码
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPageBooks(pageNo) // 调用bookdao中获取带分页的图书的函数
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, page)
}
