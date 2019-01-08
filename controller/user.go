package controller

import (
	. "driversystem/http"
	"driversystem/logic"
	"driversystem/middleware"
	"driversystem/model"
	"github.com/labstack/echo"
)

type UserController struct{}

func (this UserController) RegisterRoute(g *echo.Group) {
	userG := g.Group("/user")

	userG.POST("/register", this.Register)
	userG.POST("/login", this.Login)
	userG.POST("/change_password", this.ChangePassword, middleware.NeedLogin)
}

func (UserController) Register(ctx echo.Context) error {
	username := ctx.FormValue("username")

	if !logic.DefaultUser.CheckUsername(username) {
		return Fail(ctx, 0, "用户名不能为空")
	}

	password := ctx.FormValue("password")
	if !logic.DefaultUser.CheckPassword(password) {
		return Fail(ctx, 0, "密码不能为空")
	}

	user := logic.DefaultUser.CreateUser(username, password)
	if user == nil {
		return Fail(ctx, 0, "注册失败")
	}

	return Success(ctx, nil)
}

func (UserController) Login(ctx echo.Context) error {
	username := ctx.FormValue("username")
	if !logic.DefaultUser.CheckUsername(username) {
		return Fail(ctx, 0, "用户名不能为空")
	}

	password := ctx.FormValue("password")
	if !logic.DefaultUser.CheckPassword(password) {
		return Fail(ctx, 0, "密码不能为空")
	}

	user := logic.DefaultUser.FindOneByUsernameAndPassed(username, password)
	if user == nil {
		return Fail(ctx, 0, "登录失败")
	}

	data := map[string]interface{}{
		"uid":      user.Uid,
		"username": user.Username,
	}

	logic.SetSessionValues(ctx, data)

	return Success(ctx, data)
}

func (UserController) ChangePassword(ctx echo.Context) error {
	password := ctx.FormValue("password")
	if !logic.DefaultUser.CheckPassword(password) {
		return Fail(ctx, 0, "密码不能为空")
	}

	uid, ok := logic.GetSessionValue(ctx, "uid").(int)
	if !ok {
		return Fail(ctx, 0, "session错误")
	}

	user := &model.User{Uid: uid, Password: password}

	err := logic.DefaultUser.ChangePassword(user)
	if err != nil {
		return Fail(ctx, 0, "修改密码错误")
	}

	return Success(ctx, nil)
}
