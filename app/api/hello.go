/*
@Time : 2021/12/7 5:40 下午
@Author : fushisanlang
@File : hello.go
@Software: GoLand
*/
package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func Khronos(r *ghttp.Request) {
	g.Log().Println("访问首页")
	r.Response.WriteJson(g.Map{
		"code":    "200",
		"Message": "Khronos",
	})
}
