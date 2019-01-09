package controller

import (
	. "driversystem/http"
	"driversystem/logic"
	"driversystem/middleware"
	"driversystem/model"
	"github.com/labstack/echo"
	"strconv"
)

type CourseController struct{}

func (this CourseController) RegisterRoute(g *echo.Group) {
	courseG := g.Group("/course", middleware.NeedLogin)

	courseG.GET("/search", this.SearchCourse)
	courseG.POST("/add", this.AddCourse)
	courseG.POST("/update", this.UpdateCourse)
	courseG.POST("/delete", this.DeleteCourse)
}

func (this CourseController) SearchCourse(ctx echo.Context) error {
	//只能查询全部
	courses := logic.DefaultCourse.GetCourses()
	count := logic.DefaultCourse.GetCourseCount()

	data := map[string]interface{}{
		"courses":     courses,
		"total_count": count,
	}
	return Success(ctx, data)
}

func (this CourseController) AddCourse(ctx echo.Context) error {
	cname := ctx.FormValue("cname")
	if cname == "" {
		return Fail(ctx, 0, "没有cname")
	}

	strBeforeCour := ctx.FormValue("before_cour")
	if strBeforeCour == "" {
		return Fail(ctx, 0, "没有before_cour")
	}
	beforeCour, err := strconv.ParseInt(strBeforeCour, 10, 64)
	if err != nil {
		return Fail(ctx, 0, "before_cour错误")
	}

	course := &model.Course{
		Cname:      cname,
		BeforeCour: int(beforeCour),
	}
	err = logic.DefaultCourse.CreateCourse(course)
	if err != nil {
		return Fail(ctx, 0, "添加学科信息失败")
	}
	return Success(ctx, nil)
}

func (this CourseController) UpdateCourse(ctx echo.Context) error {
	//可以更新cname、before_cour字段
	strCno := ctx.FormValue("cno")
	if strCno == "" {
		return Fail(ctx, 0, "没有cno")
	}
	cno, err := strconv.ParseInt(strCno, 10, 64)
	if err != nil {
		return Fail(ctx, 0, "cno错误")
	}

	course := &model.Course{
		Cno: int(cno),
	}
	cols := []string{}

	cname := ctx.FormValue("cname")
	if cname != "" {
		course.Cname = cname
		cols = append(cols, "cname")
	}

	strBeforeCour := ctx.FormValue("before_cour")
	if strBeforeCour != "" {
		beforeCour, err := strconv.ParseInt(strBeforeCour, 10, 64)
		if err != nil {
			return Fail(ctx, 0, "before_cour错误")
		}
		course.BeforeCour = int(beforeCour)
		cols = append(cols, "before_cour")
	}

	ok := logic.DefaultCourse.UpdateCourse(course, cols)
	if !ok {
		return Fail(ctx, 0, "更新学科信息失败")
	}
	return Success(ctx, nil)
}

func (this CourseController) DeleteCourse(ctx echo.Context) error {
	strCno := ctx.FormValue("cno")
	if strCno == "" {
		return Fail(ctx, 0, "没有cno")
	}
	cno, err := strconv.ParseInt(strCno, 10, 64)
	if err != nil {
		return Fail(ctx, 0, "cno错误")
	}

	ok := logic.DefaultCourse.DeleteCourse(int(cno))
	if !ok {
		return Fail(ctx, 0, "删除学科信息失败")
	}
	return Success(ctx, nil)
}
