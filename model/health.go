package model

var DifferentiateMap = map[string]struct{}{
	"正常": {},
	"色弱": {},
	"色盲": {},
}

var LeftEarMap = map[string]struct{}{
	"正常": {},
	"偏弱": {},
}

var RightEarMap = map[string]struct{}{
	"正常": {},
	"偏弱": {},
}

var LegsMap = map[string]struct{}{
	"正常":  {},
	"不相等": {},
}

var PressureMap = map[string]struct{}{
	"正常": {},
	"偏高": {},
	"偏低": {},
}

type Health struct {
	Id     int     `json:"id" xorm:"pk autoincr"`
	Sno    int     `json:"sno"`
	Sname  string  `json:"sname"`
	Height float64 `json:"height" xorm:"NULL"`
	Weight float64 `json:"weight" xorm:"NULL"`
	//辩色
	Differentiate string  `json:"differentiate" xorm:"ENUM('正常','色弱','色盲') NULL"`
	LeftSight     float64 `json:"left_sight" xorm:"NULL"`
	RightSight    float64 `json:"right_sight" xorm:"NULL"`
	LeftEar       string  `json:"left_ear" xorm:"ENUM('正常','偏弱') NULL"`
	RightEar      string  `json:"right_ear" xorm:"ENUM('正常','偏弱') NULL"`
	//腿长是否相等
	Legs string `json:"legs" xorm:"ENUM('正常','不相等') NULL"`
	//血压
	Pressure string `json:"pressure" xorm:"ENUM('正常','偏高','偏低') NULL"`
	//病史
	History string `json:"history" xorm:"NULL"`
	HText   string `json:"h_text" xorm:"TEXT NULL"`
}
