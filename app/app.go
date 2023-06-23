/*
@Time : 2021/8/3 1:19 下午
@Author : fushisanlang
@File : app.go
@Software: GoLand
*/
package app

import (
	"Khronos/app/api"
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gsession"
)

func Run() {
	//配置文件
	g.Cfg().SetFileName("config.toml")
	//服务定义
	s := g.Server()
	//session相关
	s.SetConfigWithMap(g.Map{
		"SessionMaxAge":  time.Hour * 2400,
		"SessionStorage": gsession.NewStorageRedis(g.Redis()),
	})

	//组定义 user 用户相关
	group := s.Group("/khronos/user")
	//创建用户
	group.ALL("/add", api.AddUser)
	//登录
	group.ALL("/login", api.LoginUser)
	//组定义 plan 计划相关
	group = s.Group("/khronos/plan")

	group.ALL("/get", api.GetPlan)
	group.ALL("/add", api.AddPlan)
	group.ALL("/del", api.DelPlan)
	group.ALL("/done", api.DonePlan)
	group.ALL("/change", api.ChangePlan)
	group.ALL("/getallplanprocessbetween", api.GetAllPlanProcessBetween)

	
	//index
	s.BindHandler("/khronos", api.Khronos)
	//启动web
	s.Run()
}
