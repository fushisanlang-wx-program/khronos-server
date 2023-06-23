/*
@Time : 2021/12/7 5:44 下午
@Author : fushisanlang
@File : plan
@Software: GoLand
*/
package api

import (
	"Khronos/app/model"
	"Khronos/app/service"
	"github.com/gogf/gf/frame/g"
//	"fmt"
	"github.com/gogf/gf/net/ghttp"
)

func AddPlan(r *ghttp.Request) {
	//验证session，并获取用户id
	status, uid := service.GetSessionUserId(r)
	//session验证成功
	if status == true {
		var NewPlan *model.PlanInfo
		r.Parse(&NewPlan)

		if NewPlan.PlanName == "" || NewPlan.Measure == "" || NewPlan.PlanAll == 0 {
			r.Response.WriteJson(g.Map{
				"Message": "输入信息不全", "Code": 404,
			})
		} else if NewPlan.PlanAll <= 0 || NewPlan.Start < 1672502400 || NewPlan.Stop <= 1672502400 || NewPlan.PlanDone != 0 {
			r.Response.WriteJson(g.Map{
				"Message": "输入信息异常", "Code": 405,
			})
		} else {
			NewPlan.UserId = uid
			message, code := service.AddPlanByUserId(NewPlan)
			r.Response.WriteJson(g.Map{
				"Message": message, "Code": code,
			})
		}

	} else {
		r.Response.WriteJson(g.Map{
			"Message": "请先登录", "Code": 401,
		})
	}

}

//GetAllPlanProcessBetween

func GetAllPlanProcessBetween(r *ghttp.Request) {
	//验证session，并获取用户id
	status, uid := service.GetSessionUserId(r)
	//session验证成功
	if status == true {
		//var TimeBetween *model.TimeBetween
		//r.Parse(&TimeBetween)
		//1640966400 2022

		//1672502400 2023
		//if TimeBetween.Start == 0  {
		//	TimeBetween.Start=1640966400
//
//}
		//if TimeBetween.Stop == 0  {
		//	TimeBetween.Stop=1672502400

		//}
		var Start=1640966400
		var Stop=1672502400
		//fmt.Println(uid,TimeBetween.Start,TimeBetween.Stop)
		GetAllPlanProcessBetween := service.GetAllPlanProcessBetween(uid,Start,Stop)

		r.Response.WriteJson(GetAllPlanProcessBetween)

		

	} else {
		r.Response.WriteJson(g.Map{
			"Message": "请先登录", "Code": 401,
		})
	}

}


func GetPlan(r *ghttp.Request) {
	//验证session，并获取用户id
	status, uid := service.GetSessionUserId(r)
	//session验证成功
	if status == true {
		PlanList := service.GetAllPlanByUserId(uid)

		r.Response.WriteJson(PlanList)

	} else {
		r.Response.WriteJson(g.Map{
			"Message": "请先登录", "Code": 401,
		})
	}
}

func DelPlan(r *ghttp.Request) {
	//验证session，并获取用户id
	planName := r.Get("planname")
	status, uid := service.GetSessionUserId(r)
	//session验证成功
	if status == true {

		if planName == "" {
			r.Response.WriteJson(g.Map{
				"Message": "输入信息不全", "Code": 404,
			})
		} else {
			message, code := service.DelPlanByPlanNameAndUserId(planName, uid)

			r.Response.WriteJson(g.Map{
				"Message": message, "Code": code,
			})
		}

	} else {
		r.Response.WriteJson(g.Map{
			"Message": "请先登录", "Code": 401,
		})
	}
}

func DonePlan(r *ghttp.Request) {
	//验证session，并获取用户id
	planName := r.Get("planname")
	planDone := r.Get("plandone")
	status, uid := service.GetSessionUserId(r)
	//session验证成功
	if status == true {
		if planName == "" || planDone == "" {
			r.Response.WriteJson(g.Map{
				"Message": "输入信息不全", "Code": 404,
			})
		} else {

			message, code := service.DonePlanByPlanNameAndUserId(planName, planDone, uid)

			r.Response.WriteJson(g.Map{
				"Message": message, "Code": code,
			})
		}

	} else {
		r.Response.WriteJson(g.Map{
			"Message": "请先登录", "Code": 401,
		})
	}
}

func ChangePlan(r *ghttp.Request) {
	//验证session，并获取用户id
	status, uid := service.GetSessionUserId(r)
	//session验证成功
	if status == true {
		var NewPlan *model.PlanInfo
		r.Parse(&NewPlan)

		if NewPlan.PlanName == "" || NewPlan.Measure == "" || NewPlan.PlanAll == 0 {
			r.Response.WriteJson(g.Map{
				"Message": "输入信息不全", "Code": 404,
			})
		} else if NewPlan.PlanAll <= 0 || NewPlan.Start >= NewPlan.Stop || NewPlan.PlanDone != 0 {
			r.Response.WriteJson(g.Map{
				"Message": "输入信息异常", "Code": 405,
			})
		} else {
			NewPlan.UserId = uid
			message, code := service.ChangePlanByUserId(NewPlan)
			r.Response.WriteJson(g.Map{
				"Message": message, "Code": code,
			})
		}

	} else {
		r.Response.WriteJson(g.Map{
			"Message": "请先登录", "Code": 401,
		})
	}

}
