package model

import "time"

//驾照
type License struct {
	Id    int    `json:"id" xorm:"pk autoincr"`
	Sno   int    `json:"sno"`
	Sname string `json:"sname"`
	//驾驶证号
	Lno string `json:"lno"`
	//领证时间
	ReceiveTime time.Time `json:"receive_time" xorm:"DATE"`
	//领证人
	ReceiveName string `json:"receive_name"`
	LText       string `json:"l_text" xorm:"TEXT"`
}

type LicenseInfo struct {
	Id    int    `json:"id"`
	Sno   int    `json:"sno"`
	Sname string `json:"sname"`
	//驾驶证号
	Lno string `json:"lno"`
	//领证时间
	ReceiveTime string `json:"receive_time"`
	//领证人
	ReceiveName string `json:"receive_name"`
	LText       string `json:"l_text"`
}
