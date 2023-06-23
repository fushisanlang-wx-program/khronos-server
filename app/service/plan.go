/*
@Time : 2021/12/8 9:01 上午
@Author : fushisanlang
@File : plan
@Software: GoLand
*/
package service

import (
	"Khronos/app/dao"
	"Khronos/app/model"
	"github.com/gogf/gf/util/gconv"
)

func testPlanExit(NewPlan *model.PlanInfo) bool {
	planName := NewPlan.PlanName
	userId := NewPlan.UserId
	status := dao.TestPlanExit(planName, userId)
	return status
}
func DelPlanByPlanNameAndUserId(planName interface{}, userId int) (string, int) {
	PlanName := gconv.String(planName)
	status := dao.TestPlanExit(PlanName, userId)
	if status == false {
		dao.DelPlanByPlanNameAndUserId(PlanName, userId)
		message := "删除" + PlanName + "成功"
		code := 200
		return message, code
	} else {
		message := "无此计划"
		code := 404
		return message, code
	}
}
func AddPlanByUserId(NewPlan *model.PlanInfo) (string, int) {
	status := testPlanExit(NewPlan)
	if status == false {
		message := "已有同名计划"
		code := 405
		return message, code
	} else {
		dao.AddPlanByUserId(NewPlan)
		message := "添加成功"
		code := 200
		return message, code

	}
}
func ChangePlanByUserId(Plan *model.PlanInfo) (string, int) {
	status := testPlanExit(Plan)
	if status == true {
		message := "无此计划"
		code := 405
		return message, code
	} else {
		dao.ChangePlanByUserId(Plan)
		message := "调整成功"
		code := 200
		return message, code

	}
}
func GetAllPlanByUserId(userId int) []model.PlanPrint {
	plans := dao.GetAllPlanByUserId(userId)
	planList := plans.List()
	lenList := len(planList)
	var PlanInfo model.PlanPrint
	PlanList := make([]model.PlanPrint, lenList, lenList)
	for i := 0; i < lenList; i++ {
		PlanInfo.PlanName = gconv.String(planList[i]["planname"])
		PlanInfo.PlanDone = gconv.Int(planList[i]["plandone"])
		PlanInfo.PlanAll = gconv.Int(planList[i]["planall"])
		PlanInfo.Stop = gconv.Int(planList[i]["stop"])
		PlanInfo.Start = gconv.Int(planList[i]["start"])
		PlanInfo.Measure = gconv.String(planList[i]["measure"])

		PlanList[i] = PlanInfo
	}

	return PlanList
}
func GetAllPlanProcessBetween(userId ,startTiem,stopTime int) []model.AllPlanProcess {
	//fmt.Println(userId ,startTiem,stopTime)
	allPlanProcesses := dao.GetAllPlanProcessBetween(userId,startTiem,stopTime)
	allPlanProcessesList  := allPlanProcesses.List()
	lenList := len(allPlanProcessesList)
	var AllPlanProcesses model.AllPlanProcess
	AllPlanProcessesList := make([]model.AllPlanProcess, lenList, lenList)
	for i := 0; i < lenList; i++ {
		AllPlanProcesses.PlanName = gconv.String(allPlanProcessesList[i]["planname"])
		AllPlanProcesses.Process = gconv.String(allPlanProcessesList[i]["done"])
		AllPlanProcessesList[i] = AllPlanProcesses
	}

	return AllPlanProcessesList
}
func DonePlanByPlanNameAndUserId(planName interface{}, planDone interface{}, userId int) (string, int) {
	PlanName := gconv.String(planName)
	PlanDone := gconv.Int(planDone)
	if PlanDone <= 0 {
		message := "输入错误"
		code := 405
		return message, code
	} else {
		status := dao.TestPlanExit(PlanName, userId)
		if status == false {
			dao.DonePlanByPlanNameAndUserId(PlanName, PlanDone, userId)
			message := PlanName + "计划" + "记录成功"
			code := 200
			return message, code
		} else {
			message := "无此计划"
			code := 404
			return message, code
		}
	}
}
