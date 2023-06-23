/*
@Time : 2021/8/3 1:54 下午
@Author : fushisanlang
@File : user.go
@Software: GoLand
*/
package api

import (
	"Khronos/app/model"
	"Khronos/app/service"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//创建用户
func AddUser(r *ghttp.Request) {
	username := r.Get("username")
	passstr := r.Get("passstr")
	passstrre := r.Get("passstrre")
	//mail := r.Get("mail")
	if username == nil || passstr == nil || passstrre == nil {
		r.Response.WriteJson(g.Map{
			"message ": "输入信息不全", "code": 404,
		})
	} else if passstr != passstrre {
		r.Response.WriteJson(g.Map{
			"message ": "密码不同", "code": 404,
		})
	} else {
		//根据用户模板添加用户信息
		userInfo := &model.AddUser{
			UserName:  fmt.Sprintf("%v", username),
			PassStr:   service.HashPass(fmt.Sprintf("%v", passstr)),
			PassStrRe: fmt.Sprintf("%v", passstrre),
		}
		//查询同名用户
		a := service.UserNameGet(fmt.Sprintf("%v", username))
		if a != 0 {
			r.Response.WriteJson(g.Map{
				"Message": "用户存在", "code": 405,
			})
		} else {
			//创建用户
			service.AddUser(userInfo)
			r.Response.WriteJson(g.Map{
				"Message": "用户注册成功",
				"name":    username,
				"code":    200,
			})
		}
	}
}

//登录
func LoginUser(r *ghttp.Request) {
	username := r.Get("username")
	passstr := r.Get("passstr")
	if username == nil || passstr == nil {
		r.Response.WriteJson(g.Map{
			"Message": "输入信息不全", "code": 404,
		})
		//判断密码
	} else if service.PassDiff(fmt.Sprintf("%v", username), fmt.Sprintf("%v", passstr)) != true {
		
		r.Response.WriteJson(g.Map{
			"Message": "账号或密码错误", "code": 401,
		})
	} else {

		//设置session-username
		r.Session.Set("Username", fmt.Sprintf("%v", username))
		//根据username获取userid
		UserId := service.UidGetByName(fmt.Sprintf("%v", username))
		//设置session-uid
		r.Session.Set("UserId", UserId)
		r.Response.WriteJson(g.Map{
			"Message": "登录成功", "code": 200,
		})
	}
}
