package dao

import (
	"fmt"
	"strconv"

	"github.com/shuwenhe/shuwen-bookstore/utils"

	"github.com/shuwenhe/shuwen-bookstore/model"
)

// AddBook 向数据库中添加一本图书
func AddBook(b *model.Book) error {
	sql := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sql, &b.Title, &b.Author, &b.Price, &b.Sales, &b.Stock, &b.ImgPath)
	fmt.Println("b = ", b)
	if err != nil {
		return err
	}
	return nil
}

func GetBooks() ([]*model.Book, error) {
	sql := "select id,title,author,price,sales,stock,img_path from books"
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	return books, nil
}

func DeleteBook(bookId string) error {
	sql := "delete from books where id=?"
	_, err := utils.Db.Exec(sql, bookId)
	if err != nil {
		return err
	}
	return nil
}

// GetBookByID 根据图书的id从数据库中查询一本图书
func GetBookByID(bookID string) (*model.Book, error) {
	sql := "select id,title,author,price,sales,stock,img_path from books where id=?" // sql
	row := utils.Db.QueryRow(sql, bookID)                                            // 执行
	book := &model.Book{}
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath) // 为book中的字段赋值
	return book, nil
}

// UpdateBook update book
func UpdateBook(b *model.Book) error {
	sql := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"    // sql
	_, err := utils.Db.Exec(sql, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ID) // 执行
	if err != nil {
		return err
	}
	return nil
}

// GetPageBooks 获取带分页的图书信息
func GetPageBooks(pageNo string) (*model.Page, error) {
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64) // 将页码转成int64类型
	sql := "select count(*) from books"            // 获取数据库中的总记录数
	var totalRecord int64                          // 总的记录数
	row := utils.Db.QueryRow(sql)                  // 执行
	row.Scan(&totalRecord)
	var pageSize int64 = 4 // 每页显示4本书
	var totalPageNo int64  // 总页数
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	// 获取当前页中的图书
	sql2 := "select id,title,author,price,sales,stock,img_path from books limit ?,?"
	rows, err := utils.Db.Query(sql2, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book) // 将book添加到books切片中
	}
	// 创建page
	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}

// GetPageBooksByPrice 获取带分页和价格范围的图书信息
func GetPageBooksByPrice(pageNo, minPrice, maxPrice string) (*model.Page, error) {
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)                  // 将页码转成int64类型
	sql := "select count(*) from books where price between ? and ?" // 获取数据库中的总记录数
	var totalRecord int64                                           // 总的记录数
	row := utils.Db.QueryRow(sql, minPrice, maxPrice)               // 执行
	row.Scan(&totalRecord)
	var pageSize int64 = 4 // 每页显示4本书
	var totalPageNo int64  // 总页数
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	sql2 := "select id,title,author,price,sales,stock,img_path from books where price between ? and ? limit ?,?" // 获取当前页中的图书
	rows, err := utils.Db.Query(sql2, minPrice, maxPrice, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book) // 将book添加到books切片中
	}
	page := &model.Page{ // 创建page
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}
