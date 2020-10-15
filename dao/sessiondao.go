package dao

import (
	"net/http"

	"github.com/shuwenhe/shuwen-bookstore/utils"

	"github.com/shuwenhe/shuwen-bookstore/model"
)

// AddSession 添加session到数据库
func AddSession(sess *model.Session) error {
	sql := "insert into sessions values(?,?,?)"                                 // sql
	_, err := utils.Db.Exec(sql, &sess.SessionID, &sess.UserName, &sess.UserID) // 执行
	if err != nil {
		return err
	}
	return nil
}

// DeleteSession 删除数据库中的session
func DeleteSession(sessID string) error {
	sql := "delete from sessions where session_id = ?" // sql语句
	_, err := utils.Db.Exec(sql, sessID)               // 执行
	if err != nil {
		return err
	}
	return nil
}

// GetSession 根据session的Id到数据库中查询session
func GetSession(sessID string) (*model.Session, error) {
	sql := "select session_id,username,user_id from sessions where session_id = ?" // sql语句
	Stmt, err := utils.Db.Prepare(sql)                                             // 预编译
	if err != nil {
		return nil, err
	}
	row := Stmt.QueryRow(sessID) // 执行
	sess := &model.Session{}
	row.Scan(&sess.SessionID, &sess.UserName, &sess.UserID) // 扫描数据库中的字段值为session的字段赋值
	return sess, nil
}

// IsLogin 判断用户是否已经登录 false 没有登录，true 已经登录
func IsLogin(r *http.Request) (bool, *model.Session) {
	cookie, _ := r.Cookie("user") // 根据Cookie的name判断获取Cookie
	if cookie != nil {
		cookieValue := cookie.Value           // 获取Cookie的value
		session, _ := GetSession(cookieValue) // 根据cookieValue去数据库中查找与之对应的session
		if session.UserID > 0 {
			return true, session // 已经登录
		}
	}
	return false, nil // 没有登录
}
