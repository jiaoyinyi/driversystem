package controller

import (
	. "driversystem/http"
	"driversystem/logic"
	"driversystem/middleware"
	"driversystem/model"
	"github.com/labstack/echo"
	"log"
	"strconv"
)

type StudentController struct{}

func (this StudentController) RegisterRoute(g *echo.Group) {
	studentG := g.Group("/student", middleware.NeedLogin)

	studentG.GET("/search", this.SearchStudent)
	studentG.POST("/add", this.AddStudent)
	studentG.POST("/update", this.UpdateStudentInfo)
	studentG.POST("/delete", this.DeleteStudent)
}

func (this StudentController) SearchStudent(ctx echo.Context) error {
	//可以通过学号、姓名、报考车型、学员状态查询
	//此四种只能单独查询
	//后三种分页请求
	strSno := ctx.QueryParam("sno")
	if strSno != "" {
		sno, err := strconv.ParseInt(strSno, 10, 64)
		if err != nil {
			return Fail(ctx, 0, "sno类型错误")
		}
		student := logic.DefaultStudent.FindOne(int(sno))
		if student == nil {
			return Fail(ctx, 0, "查询学员信息失败")
		}
		info := logic.DefaultStudent.DealStudent(student)
		data := map[string]interface{}{
			"student": info,
		}
		return Success(ctx, data)
	}

	var count int
	var students []*model.Student
	var infos []*model.StudentInfo
	if sname := ctx.QueryParam("sname"); sname != "" {
		students = logic.DefaultStudent.GetStudentsBySname(ctx, sname)
		count = logic.DefaultStudent.GetStudentCountBySname(sname)
		goto DEAL
	}

	if carType := ctx.QueryParam("car_type"); carType != "" {
		students = logic.DefaultStudent.GetStudentsByCarType(ctx, carType)
		count = logic.DefaultStudent.GetStudentCountByCarType(carType)
		goto DEAL
	}

	if scondition := ctx.QueryParam("scondition"); scondition != "" {

		ok := logic.DefaultStudent.CheckScondition(scondition)
		if !ok {
			return Fail(ctx, 0, "scondition错误")
		}
		students = logic.DefaultStudent.GetStudentsByScondition(ctx, scondition)
		count = logic.DefaultStudent.GetStudentCountByScondition(scondition)
		goto DEAL
	}

	//获取学员信息
	students = logic.DefaultStudent.GetStudents(ctx)
	count = logic.DefaultStudent.GetStudentCount()
DEAL:
	infos = logic.DefaultStudent.DealStudents(students)
	data := map[string]interface{}{
		"total_count": count,
		"students":    infos,
	}
	return Success(ctx, data)
}

func (this StudentController) AddStudent(ctx echo.Context) error {
	student := &model.Student{}

	sname := ctx.FormValue("sname")
	if sname == "" {
		return Fail(ctx, 0, "没有sname")
	}
	student.Sname = sname

	sex := ctx.FormValue("sex")
	if sex == "" {
		return Fail(ctx, 0, "没有sex")
	}
	ok := logic.DefaultStudent.CheckSex(sex)
	if !ok {
		return Fail(ctx, 0, "sex错误")
	}
	student.Sex = sex

	strAge := ctx.FormValue("age")
	if strAge != "" {
		//可以为空
		age, err := strconv.ParseInt(strAge, 10, 64)
		if err != nil {
			return Fail(ctx, 0, "age错误")
		}
		student.Age = int(age)
	}

	identify := ctx.FormValue("identify")
	if identify == "" {
		return Fail(ctx, 0, "没有identify")
	}
	student.Identify = identify

	tel := ctx.FormValue("tel")
	if tel != "" {
		//可以为空
		student.Tel = tel
	}

	carType := ctx.FormValue("car_type")
	if carType == "" {
		return Fail(ctx, 0, "没有car_type")
	}
	student.CarType = carType

	strEnrollTime := ctx.FormValue("enroll_time")
	if strEnrollTime == "" {
		return Fail(ctx, 0, "没有enroll_time")
	}
	enrollTime, err := logic.ParseTime(strEnrollTime)
	if err != nil {
		return Fail(ctx, 0, "enroll_time错误")
	}
	student.EnrollTime = enrollTime

	strLeaveTime := ctx.FormValue("leave_time")
	if strLeaveTime != "" {
		//可以为空
		leaveTime, err := logic.ParseTime(strLeaveTime)
		if err != nil {
			return Fail(ctx, 0, "leave_time错误")
		}
		student.LeaveTime = leaveTime
	}

	scondition := ctx.FormValue("scondition")
	if scondition == "" {
		return Fail(ctx, 0, "没有scondition")
	}
	ok = logic.DefaultStudent.CheckScondition(scondition)
	if !ok {
		return Fail(ctx, 0, "scondition错误")
	}
	student.Scondition = scondition

	sText := ctx.FormValue("s_text")
	if sText != "" {
		//可以为空
		student.SText = sText
	}

	err = logic.DefaultStudent.CreateStudent(student)
	if err != nil {
		return Fail(ctx, 0, "添加学员信息失败")
	}
	return Success(ctx, nil)
}

func (this StudentController) UpdateStudentInfo(ctx echo.Context) error {
	//除了sno不可修改，其他信息都可以修改
	strSno := ctx.FormValue("sno")
	if strSno == "" {
		return Fail(ctx, 0, "没有sno")
	}
	sno, err := strconv.ParseInt(strSno, 10, 64)
	if err != nil {
		return Fail(ctx, 0, "sno错误")
	}

	log.Println(sno)
	//查找是否有该学员
	student := logic.DefaultStudent.FindOne(int(sno))
	if student == nil {
		return Fail(ctx, 0, "没有该学员信息")
	}
	log.Println(student)

	sname := ctx.FormValue("sname")
	if sname != "" {
		student.Sname = sname
	}

	sex := ctx.FormValue("sex")
	if sex != "" {
		ok := logic.DefaultStudent.CheckSex(sex)
		if !ok {
			return Fail(ctx, 0, "sex错误")
		}
		student.Sex = sex
	}

	strAge := ctx.FormValue("age")
	if strAge != "" {
		age, err := strconv.ParseInt(strAge, 10, 64)
		if err != nil {
			return Fail(ctx, 0, "age错误")
		}
		student.Age = int(age)
	}

	identify := ctx.FormValue("identify")
	if identify != "" {
		student.Identify = identify
	}

	tel := ctx.FormValue("tel")
	if tel != "" {
		student.Tel = tel
	}

	carType := ctx.FormValue("car_type")
	if carType != "" {
		student.CarType = carType
	}

	strEnrollTime := ctx.FormValue("enroll_time")
	if strEnrollTime != "" {
		enrollTime, err := logic.ParseTime(strEnrollTime)
		if err != nil {
			return Fail(ctx, 0, "enroll_time错误")
		}
		student.EnrollTime = enrollTime
	}

	strLeaveTime := ctx.FormValue("leave_time")
	if strLeaveTime != "" {
		//可以为空
		leaveTime, err := logic.ParseTime(strLeaveTime)
		if err != nil {
			return Fail(ctx, 0, "leave_time错误")
		}
		student.LeaveTime = leaveTime
	}

	scondition := ctx.FormValue("scondition")
	if scondition != "" {
		ok := logic.DefaultStudent.CheckScondition(scondition)
		if !ok {
			return Fail(ctx, 0, "scondition错误")
		}
		student.Scondition = scondition
	}

	sText := ctx.FormValue("s_text")
	if sText != "" {
		student.SText = sText
	}

	ok := logic.DefaultStudent.UpdateStudent(student)
	if !ok {
		return Fail(ctx, 0, "更新学员信息失败")
	}
	return Success(ctx, nil)
}

func (this StudentController) DeleteStudent(ctx echo.Context) error {

	strSno := ctx.FormValue("sno")
	if strSno == "" {
		return Fail(ctx, 0, "没有sno")
	}
	sno, err := strconv.ParseInt(strSno, 10, 64)
	if err != nil {
		return Fail(ctx, 0, "sno错误")
	}

	ok := logic.DefaultStudent.DeleteStudent(int(sno))
	if !ok {
		return Fail(ctx, 0, "删除学员信息失败")
	}
	return Success(ctx, nil)
}
