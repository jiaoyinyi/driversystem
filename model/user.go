package model

//管理员
type User struct {
	Uid      int    `json:"uid" xorm:"pk autoincr"`
	Username string `json:"username"`
	Password string `json:"password"`
}
