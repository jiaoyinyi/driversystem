package route

import (
	"driversystem/controller"
	"github.com/labstack/echo"
)

func RegisterAllRoutes(g *echo.Group) {
	//注册用户路由
	new(controller.UserController).RegisterRoute(g)
	//注册学员信息路由
	new(controller.StudentController).RegisterRoute(g)
	//注册学员体检信息路由
	new(controller.HealthController).RegisterRoute(g)
}
