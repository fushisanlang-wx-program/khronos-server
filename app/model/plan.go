/*
@Time : 2021/12/8 9:30 上午
@Author : fushisanlang
@File : plan
@Software: GoLand
*/
package model

type PlanPrint struct {
	PlanName string

	PlanDone int
	PlanAll  int
	Measure  string
	Start    int
	Stop     int
}
type AllPlanProcess struct {
	PlanName string
	Process string
}

type TimeBetween struct {
	Start int
	Stop int
}
type PlanInfo struct {
	UserId   int
	PlanName string
	Planid   int
	PlanDone int
	PlanAll  int
	Measure  string
	Start    int
	Stop     int
}
