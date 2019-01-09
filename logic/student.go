package logic

import (
	. "driversystem/db"
	"driversystem/model"
	"errors"
	"github.com/labstack/echo"
)

type StudentLogic struct{}

var DefaultStudent = StudentLogic{}

//插入学员信息
func (this StudentLogic) CreateStudent(s *model.Student, cols []string) error {
	if this.StudentExist(s.Sno) {
		return errors.New("student exist error")
	}

	_, err := MasterDB.Cols(cols...).Insert(s)
	return err
}

//获取学员信息 分页
func (this StudentLogic) GetStudents(ctx echo.Context) []*model.Student {
	offset, limit := DefaultPage.GetPage(ctx)

	students := make([]*model.Student, 0)
	_ = MasterDB.Limit(limit, offset).Find(&students)
	return students
}

func (this StudentLogic) GetStudentCount() int {
	s := &model.Student{}
	count, _ := MasterDB.Count(s)
	return int(count)
}

//根据学员名字获取学员信息 分页
func (this StudentLogic) GetStudentsBySname(ctx echo.Context, sname string) []*model.Student {
	offset, limit := DefaultPage.GetPage(ctx)

	students := make([]*model.Student, 0)
	_ = MasterDB.Where("sname=?", sname).Limit(limit, offset).Find(&students)
	return students
}

func (this StudentLogic) GetStudentCountBySname(sname string) int {
	s := &model.Student{}
	count, _ := MasterDB.Where("sname=?", sname).Count(s)
	return int(count)
}

//根据报考车型获取学员信息 分页
func (this StudentLogic) GetStudentsByCarType(ctx echo.Context, carType string) []*model.Student {
	offset, limit := DefaultPage.GetPage(ctx)

	students := make([]*model.Student, 0)
	_ = MasterDB.Where("car_type=?", carType).Limit(limit, offset).Find(&students)
	return students
}

func (this StudentLogic) GetStudentCountByCarType(carType string) int {
	s := &model.Student{}
	count, _ := MasterDB.Where("car_type=?", carType).Count(s)
	return int(count)
}

//根据学业状态获取学员信息 分页
func (this StudentLogic) GetStudentsByScondition(ctx echo.Context, scondition string) []*model.Student {
	offset, limit := DefaultPage.GetPage(ctx)

	students := make([]*model.Student, 0)
	_ = MasterDB.Where("scondition=?", scondition).Limit(limit, offset).Find(&students)
	return students
}

func (this StudentLogic) GetStudentCountByScondition(scondition string) int {
	s := &model.Student{}
	count, _ := MasterDB.Where("scondition=?", scondition).Count(s)
	return int(count)
}

//获取一个学员信息
func (this StudentLogic) FindOne(sno int) *model.Student {
	s := &model.Student{}
	ok, _ := MasterDB.Where("sno = ?", sno).Get(s)
	if !ok {
		return nil
	}
	return s
}

//删除学员信息
func (this StudentLogic) DeleteStudent(sno int) bool {
	s := &model.Student{}
	num, _ := MasterDB.Where("sno=?", sno).Delete(s)
	if num == 0 {
		return false
	}
	return true
}

func (this StudentLogic) UpdateStudent(s *model.Student, cols []string) bool {
	num, _ := MasterDB.Where("sno=?", s.Sno).Cols(cols...).Update(s)
	if num == 0 {
		return false
	}
	return true
}

//根据学员sno判断是否存在该学员
func (this StudentLogic) StudentExist(sno int) bool {
	exist, _ := MasterDB.Where("sno = ?", sno).Exist(new(model.Student))
	return exist
}

//检测请求过来的sex是否存在
func (this StudentLogic) CheckSex(sex string) bool {
	_, ok := model.SexMap[sex]
	return ok
}

//检测请求过来的scondition是否存在
func (this StudentLogic) CheckScondition(scondition string) bool {
	_, ok := model.SconditionMap[scondition]
	return ok
}

//处理student数据
func (this StudentLogic) DealStudent(s *model.Student) *model.StudentInfo {
	info := &model.StudentInfo{
		Sno:        s.Sno,
		Sname:      s.Sname,
		Sex:        s.Sex,
		Age:        s.Age,
		Identify:   s.Identify,
		Tel:        s.Tel,
		CarType:    s.CarType,
		EnrollTime: FormatTime(s.EnrollTime),
		LeaveTime:  FormatTime(s.LeaveTime),
		Scondition: s.Scondition,
		SText:      s.SText,
	}
	return info
}

//处理student数据
func (this StudentLogic) DealStudents(ss []*model.Student) []*model.StudentInfo {
	infos := make([]*model.StudentInfo, len(ss))
	for index, s := range ss {
		info := this.DealStudent(s)
		infos[index] = info
	}
	return infos
}
