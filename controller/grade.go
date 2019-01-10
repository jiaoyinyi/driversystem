package controller

import (
	. "driversystem/http"
	"driversystem/logic"
	"driversystem/middleware"
	"driversystem/model"
	"github.com/labstack/echo"
	"strconv"
)

type GradeController struct{}

func (this GradeController) RegisterRoute(g *echo.Group) {
	gradeG := g.Group("/grade", middleware.NeedLogin)

	gradeG.GET("/search", this.SearchGrade)
	gradeG.POST("/add", this.AddGrade)
	gradeG.POST("/update", this.UpdateGrade)
	gradeG.POST("/delete", this.DeleteGrade)
}

func (this GradeController) SearchGrade(ctx echo.Context) error {
	//根据学号、课程号查询 学号和课程号联合查询
	haveSno := false
	haveCno := false

	var sno int64
	var cno int64
	var err error

	strSno := ctx.QueryParam("sno")
	if strSno != "" {
		sno, err = strconv.ParseInt(strSno, 10, 64)
		if err != nil {
			return Fail(ctx, 0, "sno错误")
		}
		haveSno = true
	}

	strCno := ctx.QueryParam("cno")
	if strCno != "" {
		cno, err = strconv.ParseInt(strCno, 10, 64)
		if err != nil {
			return Fail(ctx, 0, "cno错误")
		}
		haveCno = true
	}

	var grades []*model.Grade
	var count int
	if haveSno && haveCno {
		grade := logic.DefaultGrade.GetGradeInfoBySnoAndCno(int(sno), int(cno))
		if grade == nil {
			return Fail(ctx, 0, "查询学员成绩失败")
		}
		info := logic.DefaultGrade.DealGrade(grade)
		data := map[string]interface{}{
			"grade_info": info,
		}
		return Success(ctx, data)
	}
	if haveSno {
		grades = logic.DefaultGrade.GetGradeInfoBySno(ctx, int(sno))
		count = logic.DefaultGrade.GetGradeInfoCountBySno(int(sno))
		goto DEAL
	}
	if haveCno {
		grades = logic.DefaultGrade.GetGradeInfoByCno(ctx, int(cno))
		count = logic.DefaultGrade.GetGradeInfoCountByCno(int(cno))
		goto DEAL
	}

	grades = logic.DefaultGrade.GetGradeInfos(ctx)
	count = logic.DefaultGrade.GetGradeInfoCount()

DEAL:
	infos := logic.DefaultGrade.DealGrades(grades)
	data := map[string]interface{}{
		"total_count": count,
		"per_page":    logic.PerPage,
		"grade_infos": infos,
	}
	return Success(ctx, data)
}

func (this GradeController) AddGrade(ctx echo.Context) error {
	grade := &model.Grade{}

	strSno := ctx.FormValue("sno")
	if strSno == "" {
		return Fail(ctx, 0, "没有sno")
	}
	sno, err := strconv.ParseInt(strSno, 10, 64)
	if err != nil {
		return Fail(ctx, 0, "sno错误")
	}
	exist := logic.DefaultStudent.StudentExist(int(sno))
	if !exist {
		return Fail(ctx, 0, "没有该学员")
	}
	grade.Sno = int(sno)

	strCno := ctx.FormValue("cno")
	if strCno == "" {
		return Fail(ctx, 0, "没有cno")
	}
	cno, err := strconv.ParseInt(strCno, 10, 64)
	if err != nil {
		return Fail(ctx, 0, "cno错误")
	}
	exist = logic.DefaultCourse.CourseExist(int(cno))
	if !exist {
		return Fail(ctx, 0, "没有该科目")
	}
	grade.Cno = int(cno)
	//最近的一次考试时间
	strLastTime := ctx.FormValue("last_time")
	if strLastTime == "" {
		return Fail(ctx, 0, "没有last_time")
	}
	lastTime, err := logic.ParseTime(strLastTime)
	if err != nil {
		return Fail(ctx, 0, "last_time错误")
	}
	grade.LastTime = lastTime
	//考了多少次 默认一次
	strTimes := ctx.FormValue("times")
	if strTimes != "" {
		times, err := strconv.ParseInt(strTimes, 10, 64)
		if err != nil {
			return Fail(ctx, 0, "times错误")
		}
		grade.Times = int(times)
	} else {
		grade.Times = 1
	}
	//成绩
	strG := ctx.FormValue("grade")
	if strG == "" {
		return Fail(ctx, 0, "没有grade")
	}
	g, err := strconv.ParseFloat(strG, 64)
	if err != nil {
		return Fail(ctx, 0, "grade错误")
	}
	grade.Grade = g

	err = logic.DefaultGrade.CreateGrade(grade)
	if err != nil {
		return Fail(ctx, 0, "添加学员科目成绩失败")
	}
	return Success(ctx, nil)
}

func (this GradeController) UpdateGrade(ctx echo.Context) error {
	//只能更新last_time、times、grade字段
	strId := ctx.FormValue("id")
	if strId == "" {
		return Fail(ctx, 0, "没有id")
	}
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		return Fail(ctx, 0, "id错误")
	}
	grade := &model.Grade{Id: int(id)}
	cols := make([]string, 0)

	strLastTime := ctx.FormValue("last_time")
	if strLastTime != "" {
		lastTime, err := logic.ParseTime(strLastTime)
		if err != nil {
			return Fail(ctx, 0, "last_time错误")
		}
		grade.LastTime = lastTime
		cols = append(cols, "last_time")
	}

	strTimes := ctx.FormValue("times")
	if strTimes != "" {
		times, err := strconv.ParseInt(strTimes, 10, 64)
		if err != nil {
			return Fail(ctx, 0, "times错误")
		}
		grade.Times = int(times)
		cols = append(cols, "times")
	}

	//成绩
	strG := ctx.FormValue("grade")
	if strG != "" {
		g, err := strconv.ParseFloat(strG, 64)
		if err != nil {
			return Fail(ctx, 0, "grade错误")
		}
		grade.Grade = g
		cols = append(cols, "grade")
	}

	ok := logic.DefaultGrade.UpdateGradeInfo(grade, cols)
	if !ok {
		return Fail(ctx, 0, "更新学员科目成绩失败")
	}
	return Success(ctx, nil)
}

func (this GradeController) DeleteGrade(ctx echo.Context) error {
	strId := ctx.FormValue("id")
	if strId == "" {
		return Fail(ctx, 0, "没有id")
	}
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		return Fail(ctx, 0, "id错误")
	}

	ok := logic.DefaultGrade.DeleteGradeInfo(int(id))
	if !ok {
		return Fail(ctx, 0, "删除学员科目成绩失败")
	}
	return Success(ctx, nil)
}
