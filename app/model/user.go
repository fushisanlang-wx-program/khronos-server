/*
@Time : 2021/8/3 1:58 下午
@Author : fushisanlang
@File : user
@Software: GoLand
*/
//package model
package model

//注册账户信息
type AddUser struct {
	UserName  string
	PassStr   string
	PassStrRe string
}

//密码
type PassStr struct {
	PassStr string
}

//uid
type UserId struct {
	UserId int
}
