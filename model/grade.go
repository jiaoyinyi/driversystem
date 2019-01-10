package model

import "time"

type Grade struct {
	Id  int `json:"id" xorm:"pk autoincr"`
	Sno int `json:"sno"`
	Cno int `json:"cno"`
	//考试时间
	LastTime time.Time `json:"last_time" xorm:"DATE"`
	//考试次数
	Times int     `json:"times"`
	Grade float64 `json:"grade"`
}

type GradeInfo struct {
	Id  int `json:"id"`
	Sno int `json:"sno"`
	Cno int `json:"cno"`
	//考试时间
	LastTime string `json:"last_time"`
	//考试次数
	Times int     `json:"times"`
	Grade float64 `json:"grade"`
}
