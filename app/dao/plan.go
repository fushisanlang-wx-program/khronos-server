/*
@Time : 2021/12/8 9:01 上午
@Author : fushisanlang
@File : plan
@Software: GoLand
*/
package dao

import (
	"Khronos/app/model"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	_ "github.com/mattn/go-sqlite3"
)

func GetAllPlanProcessBetween(userId, startTime, stopTime int) gdb.Result {
	AllPlanProcess, _ := g.DB().GetAll("select planname,plandone/planall*100,concat(truncate(plandone/planall*100,2),'%') as done from plan where userid = ? and start >= ? and stop < ? ORDER BY 2 desc ", userId, startTime, stopTime)
	return AllPlanProcess
}
func GetAllPlanByUserId(userId int) gdb.Result {
	plan, _ := g.DB().GetAll("select * from plan where userid = ? ", userId)
	return plan
}
func TestPlanExit(planName string, userId int) bool {
	a, _ := g.DB().Model("plan").Where("planname", planName).Where("userid", userId).Count()

	if a != 0 {
		return false
	} else {
		return true
	}
}
func AddPlanByUserId(NewPlan *model.PlanInfo) {
	g.DB().Insert("plan", gdb.Map{
		"UserId":   NewPlan.UserId,
		"PlanName": NewPlan.PlanName,
		//"Planid":   NewPlan.Planid,
		"PlanDone": NewPlan.PlanDone,
		"PlanAll":  NewPlan.PlanAll,
		"Measure":  NewPlan.Measure,
		"Start":    NewPlan.Start,
		"Stop":     NewPlan.Stop,
	})

}
func DelPlanByPlanNameAndUserId(PlanName string, UserId int) {
	g.DB().GetAll("delete from plan where userid = ? and planname = ?", UserId, PlanName)

}
func DelPlanByUserId(NewPlan *model.PlanInfo) {
	g.DB().Insert("plan", gdb.Map{
		"UserId":   NewPlan.UserId,
		"PlanName": NewPlan.PlanName,
		"Planid":   NewPlan.Planid,
		"PlanDone": NewPlan.PlanDone,
		"PlanAll":  NewPlan.PlanAll,
		"Measure":  NewPlan.Measure,
		"Start":    NewPlan.Start,
		"Stop":     NewPlan.Stop,
	})
}
func DonePlanByPlanNameAndUserId(PlanName string, PlanDone int, UserId int) {
	g.DB().GetAll("UPDATE plan SET plandone = ? WHERE planname = ? and userid = ? ;", PlanDone, PlanName, UserId)

}
func ChangePlanByUserId(Plan *model.PlanInfo) {

	g.DB().GetAll("UPDATE  plan SET start = ?, stop = ?, planall = ?, measure = ? WHERE planname = ? and userid = ? ;", Plan.Start, Plan.Stop, Plan.PlanAll, Plan.Measure, Plan.PlanName, Plan.UserId)

}
