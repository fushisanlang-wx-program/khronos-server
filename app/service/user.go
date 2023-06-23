/*
@Time : 2021/8/4 1:14 下午
@Author : fushisanlang
@File : password
@Software: GoLand
*/
package service

import (
	"Khronos/app/dao"
	"Khronos/app/model"
	"encoding/json"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"golang.org/x/crypto/bcrypt"
	"unsafe"
)

//加密密码
func HashPass(password string) string {
	passHashd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passHashd)
}

//密码校验
func PassDiff(username string, password string) bool {
	passInDb := dao.PassGetByName(username)
	//err := bcrypt.CompareHashAndPassword([]byte(passInDb), []byte(HashPass(password)))
	err := bcrypt.CompareHashAndPassword([]byte(passInDb), []byte(password))

	if err == nil {

		return true
	} else {

		return false
	}

}

//获取session中的username
func GetSessionUserName(r *ghttp.Request) (bool, interface{}) {
	a := r.Session.Map()["Username"]
	//b := fmt.Sprintf("%v", a)
	//fmt.Println(b)
	if a == nil {

		return false, ""
	} else {

		return true, a
	}
}

//获取session中的uid
func GetSessionUserId(r *ghttp.Request) (bool, int) {

	a := r.Session.Map()["UserId"]

	if a == nil {

		return false, 0
	} else {
		uid64, _ := a.(json.Number).Int64()

		UserId := *(*int)(unsafe.Pointer(&uid64))
		return true, UserId
	}

}

func AddUser(userInfo *model.AddUser) {
	g.Log().Print(gconv.String(userInfo.UserName) + ": AddUser")

	dao.UserAdd(userInfo)
}

func UserNameGet(name string) int {
	a := dao.UserNameGet(name)
	return a
}
func UidGetByName(name string) int {
	UserId := dao.UidGetByName(name)
	return UserId
}
