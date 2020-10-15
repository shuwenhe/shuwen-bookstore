package dao

import (
	"github.com/shuwenhe/shuwen-bookstore/utils"

	"github.com/shuwenhe/shuwen-bookstore/model"
)

func CheckUserNameAndPassword(username string, password string) (*model.User, error) {
	sql := "select id,username,password,email from users where username = ? and password = ?"
	row := utils.Db.QueryRow(sql, username, password)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

func CheckUserName(username string) (*model.User, error) {
	sql := "select id,username,password,email from users where username = ?"
	row := utils.Db.QueryRow(sql, username)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}

func SaveUser(username string, password string, email string) error {
	sql := "insert into users(username,password,email)values(?,?,?)"
	_, err := utils.Db.Exec(sql, username, password, email)
	if err != nil {
		return err
	}
	return nil
}
