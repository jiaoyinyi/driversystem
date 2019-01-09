package controller

import (
	. "driversystem/http"
	"driversystem/logic"
	"driversystem/middleware"
	"driversystem/model"
	"github.com/labstack/echo"
	"strconv"
)

type LicenseController struct{}

func (this LicenseController) RegisterRoute(g *echo.Group) {
	licenseG := g.Group("/license", middleware.NeedLogin)

	licenseG.GET("/search", this.SearchLicenseInfo)
	licenseG.POST("/add", this.AddLicenseInfo)
	licenseG.POST("/update", this.UpdateLicenseInfo)
	licenseG.POST("/delete", this.DeleteLicenseInfo)
}

func (this LicenseController) SearchLicenseInfo(ctx echo.Context) error {
	//可以通过学号、姓名、驾照号
	//此三种只能单独查询
	//姓名分页请求

	var lisence *model.License
	var lisences []*model.License
	var count int
	//根据学号
	if strSno := ctx.FormValue("sno"); strSno != "" {
		sno, err := strconv.ParseInt(strSno, 10, 64)
		if err != nil {
			return Fail(ctx, 0, "sno错误")
		}
		lisence = logic.DefaultLicense.FindOneBySno(int(sno))
		if lisence == nil {
			return Fail(ctx, 0, "查询学员驾照信息失败")
		}
		goto DEALONE
	}
	//根据驾驶证号
	if lno := ctx.FormValue("lno"); lno != "" {
		lisence = logic.DefaultLicense.FindOneByLno(lno)
		if lisence == nil {
			return Fail(ctx, 0, "查询学员驾照信息失败")
		}
		goto DEALONE
	}
	//根据学员名称
	if sname := ctx.FormValue("sname"); sname != "" {
		lisences = logic.DefaultLicense.GetLicenseInfoBySname(ctx, sname)
		count = logic.DefaultLicense.GetLicenseInfoCountBySname(sname)
		goto DEALMANY
	}

	lisences = logic.DefaultLicense.GetLicenseInfos(ctx)
	count = logic.DefaultLicense.GetLicenseInfoCount()
	goto DEALMANY

DEALONE:
	{
		info := logic.DefaultLicense.DealLicense(lisence)
		data := map[string]interface{}{
			"license_info": info,
		}
		return Success(ctx, data)
	}

DEALMANY:
	{
		infos := logic.DefaultLicense.DealLicenses(lisences)
		data := map[string]interface{}{
			"total_count":   count,
			"per_page":      logic.PerPage,
			"license_infos": infos,
		}
		return Success(ctx, data)
	}
}

func (this LicenseController) AddLicenseInfo(ctx echo.Context) error {
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

	license := &model.License{
		Sno:   student.Sno,
		Sname: student.Sname,
	}
	cols := []string{
		"sno", "sname", "lno",
	}

	lno := ctx.FormValue("lno")
	if lno == "" {
		return Fail(ctx, 0, "没有lno")
	}
	license.Lno = lno

	strReceiveTime := ctx.FormValue("receive_time")
	if strReceiveTime != "" {
		receiveTime, err := logic.ParseTime(strReceiveTime)
		if err != nil {
			return Fail(ctx, 0, "receive_time错误")
		}
		license.ReceiveTime = receiveTime
		cols = append(cols, "receive_time")
	}

	receiveName := ctx.FormValue("receive_name")
	if receiveName != "" {
		license.ReceiveName = receiveName
		cols = append(cols, "receive_name")
	}

	lText := ctx.FormValue("l_text")
	if lText != "" {
		license.LText = lText
		cols = append(cols, "l_text")
	}

	err = logic.DefaultLicense.CreateLicenseInfo(license, cols)
	if err != nil {
		return Fail(ctx, 0, "添加学员领证信息失败")
	}
	return Success(ctx, nil)
}

func (this LicenseController) UpdateLicenseInfo(ctx echo.Context) error {
	//只能更新lno、receive_time、receive_name、l_text字段
	strId := ctx.FormValue("id")
	if strId == "" {
		return Fail(ctx, 0, "没有id")
	}
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		return Fail(ctx, 0, "id错误")
	}
	license := logic.DefaultLicense.FindOne(int(id))
	if license == nil {
		return Fail(ctx, 0, "没有该驾照信息")
	}
	cols := []string{}

	lno := ctx.FormValue("lno")
	if lno != "" {
		license.Lno = lno
		cols = append(cols, "lno")
	}

	strReceiveTime := ctx.FormValue("receive_time")
	if strReceiveTime != "" {
		receiveTime, err := logic.ParseTime(strReceiveTime)
		if err != nil {
			return Fail(ctx, 0, "receive_time错误")
		}
		license.ReceiveTime = receiveTime
		cols = append(cols, "receive_time")
	}

	receiveName := ctx.FormValue("receive_name")
	if receiveName != "" {
		license.ReceiveName = receiveName
		cols = append(cols, "receive_name")
	}

	lText := ctx.FormValue("l_text")
	if lText != "" {
		license.LText = lText
		cols = append(cols, "l_text")
	}

	ok := logic.DefaultLicense.UpdateLicenseInfo(license, cols)
	if !ok {
		return Fail(ctx, 0, "更新学员驾照信息失败")
	}
	return Success(ctx, nil)
}

func (this LicenseController) DeleteLicenseInfo(ctx echo.Context) error {
	strId := ctx.FormValue("id")
	if strId == "" {
		return Fail(ctx, 0, "没有id")
	}
	id, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		return Fail(ctx, 0, "id错误")
	}

	ok := logic.DefaultLicense.DeleteLicenseInfo(int(id))
	if !ok {
		return Fail(ctx, 0, "删除学员驾照信息失败")
	}
	return Success(ctx, nil)
}
