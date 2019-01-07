package model

import "time"

var SexMap = map[int]string{
	0: "女",
	1: "男",
}

var SconditionMap = map[int]string{
	0: "学习",
	1: "结业",
	2: "退学",
}

//学生
type Student struct {
	Sno   int    `json:"sno" xorm:"pk"`
	Sname string `json:"sname"`
	Sex   string `json:"sex" xorm:"ENUM('女','男')"`
	Age   int    `json:"age"`
	//身份证号
	Identify string `json:"identify"`
	Tel      string `json:"tel"`
	//报考车型
	CarType string `json:"car_type"`
	//入学时间
	EnrollTime time.Time `json:"enroll_time" xorm:"DATE"`
	//毕业时间
	LeaveTime time.Time `json:"leave_time" xorm:"DATE"`
	//学业状态
	Scondition string `json:"scondition" xorm:"ENUM('学习','结业','退学')"`
	SText      string `json:"s_text" xorm:"TEXT"`
}
