package controller

import (
	. "driversystem/http"
	"driversystem/logic"
	"driversystem/middleware"
	"driversystem/model"
	"github.com/labstack/echo"
	"strconv"
)

type HealthController struct{}

func (this HealthController) RegisterRoute(g *echo.Group) {
	healthG := g.Group("/health", middleware.NeedLogin)

	healthG.GET("/search", this.SearchHealthInfo)
	healthG.POST("/add", this.AddHealthInfo)
	healthG.POST("/update", this.UpdateHealthInfo)
	healthG.POST("/delete", this.DeleteHealthInfo)
}

func (this HealthController) SearchHealthInfo(ctx echo.Context) error {
	//可以通过学号、姓名、辩色
	//此三种只能单独查询
	//后两种分页请求

	strSno := ctx.QueryParam("sno")
	if strSno != "" {
		sno, err := strconv.ParseInt(strSno, 10, 64)
		if err != nil {
			return Fail(ctx, 0, "sno类型错误")
		}
		health := logic.DefaultHealth.FindOneBySno(int(sno))
		if health == nil {
			return Fail(ctx, 0, "没有该学员")
		}
		data := map[string]interface{}{
			"health_info": health,
		}
		return Success(ctx, data)
	}

	var count int
	var healths []*model.Health
	//根据学员名字
	if sname := ctx.QueryParam("sname"); sname != "" {
		healths = logic.DefaultHealth.GetHealthInfoBySname(ctx, sname)
		count = logic.DefaultHealth.GetHealthInfoCountBySname(sname)
		goto DEAL
	}
	//根据辩色
	if differentiate := ctx.QueryParam("differentiate"); differentiate != "" {
		ok := logic.DefaultHealth.CheckDifferentiate(differentiate)
		if !ok {
			return Fail(ctx, 0, "differentiate错误")
		}
		healths = logic.DefaultHealth.GetHealthInfoByDifferentiate(ctx, differentiate)
		count = logic.DefaultHealth.GetHealthInfoCountByDifferentiate(differentiate)
		goto DEAL
	}

	//获取所有学员体检信息 分页
	healths = logic.DefaultHealth.GetHealthInfos(ctx)
	count = logic.DefaultHealth.GetHealthInfoCount()

DEAL:
	data := map[string]interface{}{
		"total_count":  count,
		"per_page":     logic.PerPage,
		"health_infos": healths,
	}
	return Success(ctx, data)
}

func (this HealthController) AddHealthInfo(ctx echo.Context) error {
	strSno := ctx.FormValue("sno")
	if strSno == "" {
		return Fail(ctx, 0, "没有sno")
	}
	sno, err := strconv.ParseInt(strSno, 10, 64)
	if err != nil {
		return Fail(ctx, 0, "sno错误")
	}

	student := logic.DefaultStudent.FindOne(int(sno))
	if student == nil {
		return Fail(ctx, 0, "没有该学员信息")
	}

	health := &model.Health{
		Sno:   student.Sno,
		Sname: student.Sname,
	}
	cols := []string{
		"sno", "sname",
	}

	strHeight := ctx.FormValue("height")
	if strHeight != "" {
		height, err := strconv.ParseFloat(strHeight, 64)
		if err != nil {
			return Fail(ctx, 0, "height错误")
		}
		health.Height = height
		cols = append(cols, "height")
	}

	strWeight := ctx.FormValue("weight")
	if strWeight != "" {
		weight, err := strconv.ParseFloat(strWeight, 64)
		if err != nil {
			return Fail(ctx, 0, "weight错误")
		}
		health.Weight = weight
		cols = append(cols, "weight")
	}

	differentiate := ctx.FormValue("differentiate")
	if differentiate != "" {
		ok := logic.DefaultHealth.CheckDifferentiate(differentiate)
		if !ok {
			return Fail(ctx, 0, "differentiate错误")
		}
		health.Differentiate = differentiate
		cols = append(cols, "differentiate")
	}

	strLeftSight := ctx.FormValue("left_sight")
	if strLeftSight != "" {
		leftSight, err := strconv.ParseFloat(strLeftSight, 64)
		if err != nil {
			return Fail(ctx, 0, "left_sight错误")
		}
		health.LeftSight = leftSight
		cols = append(cols, "left_sight")
	}

	strRightSight := ctx.FormValue("right_sight")
	if strRightSight != "" {
		rightSight, err := strconv.ParseFloat(strRightSight, 64)
		if err != nil {
			return Fail(ctx, 0, "right_sight错误")
		}
		health.RightSight = rightSight
		cols = append(cols, "right_sight")
	}

	leftEar := ctx.FormValue("left_ear")
	if leftEar != "" {
		ok := logic.DefaultHealth.CheckLeftEar(leftEar)
		if !ok {
			return Fail(ctx, 0, "left_ear错误")
		}
		health.LeftEar = leftEar
		cols = append(cols, "left_ear")
	}

	rightEar := ctx.FormValue("right_ear")
	if rightEar != "" {
		ok := logic.DefaultHealth.CheckRightEar(rightEar)
		if !ok {
			return Fail(ctx, 0, "right_ear错误")
		}
		health.RightEar = rightEar
		cols = append(cols, "right_ear")
	}

	legs := ctx.FormValue("legs")
	if legs != "" {
		ok := logic.DefaultHealth.CheckLegs(legs)
		if !ok {
			return Fail(ctx, 0, "legs错误")
		}
		health.Legs = legs
		cols = append(cols, "legs")
	}

	pressure := ctx.FormValue("pressure")
	if pressure != "" {
		ok := logic.DefaultHealth.CheckPressure(pressure)
		if !ok {
			return Fail(ctx, 0, "pressure错误")
		}
		health.Pressure = pressure
		cols = append(cols, "pressure")
	}

	history := ctx.FormValue("history")
	if history != "" {
		health.History = history
		cols = append(cols, "history")
	}

	hText := ctx.FormValue("h_text")
	if hText != "" {
		health.HText = hText
		cols = append(cols, "h_text")
	}

	err = logic.DefaultHealth.CreateHealthInfo(health, cols)
	if err != nil {
		return Fail(ctx, 0, "添加学员体检信息失败")
	}
	return Success(ctx, nil)
}

func (this HealthController) UpdateHealthInfo(ctx echo.Context) error {

	strId := ctx.FormValue("id")
	if strId == "" {
		return Fail(ctx, 0, "没有id")
	}
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		return Fail(ctx, 0, "id错误")
	}

	health := &model.Health{Id: int(id)}
	cols := []string{}

	strHeight := ctx.FormValue("height")
	if strHeight != "" {
		height, err := strconv.ParseFloat(strHeight, 64)
		if err != nil {
			return Fail(ctx, 0, "height错误")
		}
		health.Height = height
		cols = append(cols, "height")
	}

	strWeight := ctx.FormValue("weight")
	if strWeight != "" {
		weight, err := strconv.ParseFloat(strWeight, 64)
		if err != nil {
			return Fail(ctx, 0, "weight错误")
		}
		health.Weight = weight
		cols = append(cols, "weight")
	}

	differentiate := ctx.FormValue("differentiate")
	if differentiate != "" {
		ok := logic.DefaultHealth.CheckDifferentiate(differentiate)
		if !ok {
			return Fail(ctx, 0, "differentiate错误")
		}
		health.Differentiate = differentiate
		cols = append(cols, "differentiate")
	}

	strLeftSight := ctx.FormValue("left_sight")
	if strLeftSight != "" {
		leftSight, err := strconv.ParseFloat(strLeftSight, 64)
		if err != nil {
			return Fail(ctx, 0, "left_sight错误")
		}
		health.LeftSight = leftSight
		cols = append(cols, "left_sight")
	}

	strRightSight := ctx.FormValue("right_sight")
	if strRightSight != "" {
		rightSight, err := strconv.ParseFloat(strRightSight, 64)
		if err != nil {
			return Fail(ctx, 0, "right_sight错误")
		}
		health.RightSight = rightSight
		cols = append(cols, "right_sight")
	}

	leftEar := ctx.FormValue("left_ear")
	if leftEar != "" {
		ok := logic.DefaultHealth.CheckLeftEar(leftEar)
		if !ok {
			return Fail(ctx, 0, "left_ear错误")
		}
		health.LeftEar = leftEar
		cols = append(cols, "left_ear")
	}

	rightEar := ctx.FormValue("right_ear")
	if rightEar != "" {
		ok := logic.DefaultHealth.CheckRightEar(rightEar)
		if !ok {
			return Fail(ctx, 0, "right_ear错误")
		}
		health.RightEar = rightEar
		cols = append(cols, "right_ear")
	}

	legs := ctx.FormValue("legs")
	if legs != "" {
		ok := logic.DefaultHealth.CheckLegs(legs)
		if !ok {
			return Fail(ctx, 0, "legs错误")
		}
		health.Legs = legs
		cols = append(cols, "legs")
	}

	pressure := ctx.FormValue("pressure")
	if pressure != "" {
		ok := logic.DefaultHealth.CheckPressure(pressure)
		if !ok {
			return Fail(ctx, 0, "pressure错误")
		}
		health.Pressure = pressure
		cols = append(cols, "pressure")
	}

	history := ctx.FormValue("history")
	if history != "" {
		health.History = history
		cols = append(cols, "history")
	}

	hText := ctx.FormValue("h_text")
	if hText != "" {
		health.HText = hText
		cols = append(cols, "h_text")
	}

	ok := logic.DefaultHealth.UpdateHealthInfo(health, cols)
	if !ok {
		return Fail(ctx, 0, "更新学员体检信息失败")
	}
	return Success(ctx, nil)
}

func (this HealthController) DeleteHealthInfo(ctx echo.Context) error {
	strId := ctx.FormValue("id")
	if strId == "" {
		return Fail(ctx, 0, "没有id")
	}
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		return Fail(ctx, 0, "id错误")
	}

	ok := logic.DefaultHealth.DeleteHealthInfo(int(id))
	if !ok {
		return Fail(ctx, 0, "删除学员体检信息失败")
	}
	return Success(ctx, nil)
}
