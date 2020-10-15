package dao

import (
	"fmt"
	"testing"

	"github.com/shuwenhe/shuwen-bookstore/model"
)

func testAddSession(t *testing.T) {
	sess := &model.Session{
		SessionID: "15010729356",
		UserName:  "admin",
		UserID:    8,
	}
	AddSession(sess)
}

func testDeleteSession(t *testing.T) {
	DeleteSession("15010729356")
}

func testGetSession(t *testing.T) {
	sess, _ := GetSession("07def599-fbfd-406a-7521-bed09d2230ec")
	fmt.Println("sess = ", sess)
}
