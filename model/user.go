package model

//管理员
type User struct {
	Username string `json:"username" xorm:"pk"`
	Password string `json:"password"`
}
