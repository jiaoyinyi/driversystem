package logic

import (
	. "driversystem/db"
	"driversystem/model"
	"errors"
	"log"
)

type UserLogic struct{}

var DefaultUser = UserLogic{}

func (this UserLogic) CreateUser(username, password string) (*model.User) {
	user := &model.User{
		Username: username,
		Password: password,
	}

	if this.UserExists("username", username) {
		log.Println("username exist:", username)
		return nil
	}

	_, err := MasterDB.Insert(user)
	if err != nil {
		return nil
	}
	return user
}

func (this UserLogic) ChangePassword(user *model.User) error {
	u := this.FindOne(user.Uid)
	if u == nil {
		return errors.New("change password error")
	}

	u.Password = user.Password
	_, err := MasterDB.Where("uid=?", u.Uid).Update(u)
	return err
}

func (UserLogic) FindOne(uid int) *model.User {
	user := &model.User{
		Uid: uid,
	}
	ok, _ := MasterDB.Where("uid=?", uid).Get(user)
	if !ok {
		return nil
	}
	return user
}

func (UserLogic) FindOneByUsernameAndPassed(username, password string) (*model.User) {
	user := &model.User{
		Username: username,
		Password: password,
	}

	ok, _ := MasterDB.Where("username = ? And password = ?", username, password).Get(user)
	if !ok {
		return nil
	}
	return user
}

func (UserLogic) UserExists(field, val string) bool {
	user := &model.User{}
	exist, _ := MasterDB.Where(field+" = ?", val).Exist(user)
	return exist
}

func (UserLogic) CheckUsername(username string) bool {
	//目前简单验证
	if len(username) == 0 {
		return false
	}
	return true
}

func (UserLogic) CheckPassword(password string) bool {
	////目前简单验证
	if len(password) == 0 {
		return false
	}
	return true
}
