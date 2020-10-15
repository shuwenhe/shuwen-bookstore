package dao

import (
	"fmt"
	"testing"
)

func testCheckUserNameAndPassword(t *testing.T) {
	user, _ := CheckUserNameAndPassword("admin", "123456")
	fmt.Println("user = ", user)
}

func testCheckUserName(t *testing.T) {
	user, _ := CheckUserName("admin2")
	fmt.Println("user = ", user)
}

func testSaveUser(t *testing.T) {
	err := SaveUser("admin4", "123456", "admin4@qq.com")
	if err != nil {
		err.Error()
	}
}
