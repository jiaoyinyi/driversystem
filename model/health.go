package model

var DifferentiateMap = map[int]string{
	0: "正常",
	1: "色弱",
	2: "色盲",
}

var LeftEarMap = map[int]string{
	0: "正常",
	1: "偏弱",
}

var RightEarMap = map[int]string{
	0: "正常",
	1: "偏弱",
}

var LegsMap = map[int]string{
	0: "正常",
	1: "不相等",
}

var PressureMap = map[int]string{
	0: "正常",
	1: "偏高",
	2: "偏低",
}

type Health struct {
	Id     int     `json:"id" xorm:"pk autoincr"`
	Sno    int     `json:"sno"`
	Sname  string  `json:"sname"`
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
	//辩色
	Differentiate string  `json:"differentiate" xorm:"ENUM('正常','色弱','色盲')"`
	LeftSight     float64 `json:"left_sight"`
	RightSight    float64 `json:"right_sight"`
	LeftEar       string  `json:"left_ear" xorm:"ENUM('正常','偏弱')"`
	RightEar      string  `json:"right_ear" xorm:"ENUM('正常','偏弱')"`
	//腿长是否相等
	Legs string `json:"legs" xorm:"ENUM('正常','不相等')"`
	//血压
	Pressure string `json:"pressure" xorm:"ENUM('正常','偏高','偏低')"`
	//病史
	History string `json:"history"`
	HText   string `json:"h_text" xorm:"TEXT"`
}
