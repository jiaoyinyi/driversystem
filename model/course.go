package model

type Course struct {
	Cno   int    `json:"cno" xorm:"pk autoincr"`
	Cname string `json:"cname"`
	//先行考试科目
	BeforeCour int `json:"before_cour"`
}
