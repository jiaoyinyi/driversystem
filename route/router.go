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
	//注册学员驾照信息路由
	new(controller.LicenseController).RegisterRoute(g)
	//注册学科信息路由
	new(controller.CourseController).RegisterRoute(g)
	//注册学员成绩信息路由
	new(controller.GradeController).RegisterRoute(g)
}
