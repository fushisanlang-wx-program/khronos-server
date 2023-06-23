/*
@Time : 2021/12/7 4:55 下午
@Author : fushisanlang
@File : main.go
@Software: GoLand
*/
package main

import (
	"Khronos/app"
	"github.com/gogf/gf/frame/g"
)

//入口函数
func main() {
	g.Log().Println("Khronos 服务启动")

	//启动web项目
	app.Run()

}
