package dao

import (
	"github.com/shuwenhe/shuwen-bookstore/utils"

	"github.com/shuwenhe/shuwen-bookstore/model"
)

// AddOrder add order
func AddOrder(order *model.Order) error {
	sql := "insert into orders(id,create_time,total_count,total_amount,state,user_id) values (?,?,?,?,?,?)"

	_, err := utils.Db.Exec(sql, order.OrderID, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserID)
	if err != nil {
		return err
	}
	return nil
}

// GetOrders Get all orders in the database
func GetOrders() ([]*model.Order, error) {
	sql := "select id,create_time,total_count,total_amount,state,user_id from orders" // sql
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserID)
		orders = append(orders, order)
	}
	return orders, nil
}

func GetOrderByUserID(userID int) ([]*model.Order, error) {
	sql := "select id,create_time,total_count,total_amount,state,user_id from orders where user_id=?"
	rows, err := utils.Db.Query(sql, userID)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserID)
		orders = append(orders, order)
	}
	return orders, nil
}

func UpdateOrderState(orderID string, state int64) error {
	sql := "update orders set state = ? where id = ?"
	_, err := utils.Db.Exec(sql, state, orderID)
	if err != nil {
		return err
	}
	return nil
}
