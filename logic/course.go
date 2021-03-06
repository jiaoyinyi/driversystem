package logic

import (
	. "driversystem/db"
	"driversystem/model"
)

type CourseLogic struct{}

var DefaultCourse = CourseLogic{}

func (this CourseLogic) CreateCourse(c *model.Course) error {
	_, err := MasterDB.Insert(c)
	return err
}

func (this CourseLogic) GetCourses() []*model.Course {
	cs := make([]*model.Course, 0)
	MasterDB.Find(&cs)
	return cs
}

func (this CourseLogic) GetCourseCount() int {
	c := &model.Course{}
	num, _ := MasterDB.Count(c)
	return int(num)
}

func (this CourseLogic) UpdateCourse(c *model.Course, cols []string) bool {
	num, _ := MasterDB.Where("cno=?", c.Cno).Cols(cols...).Update(c)
	if num == 0 {
		return false
	}
	return true
}

func (this CourseLogic) DeleteCourse(cno int) bool {
	sess := MasterDB.NewSession()
	defer sess.Close()

	sess.Begin()
	c := &model.Course{}
	_, err := MasterDB.Where("cno=?", cno).Delete(c)
	if err != nil {
		sess.Rollback()
		return false
	}
	//当删除了该cno的数据，要连同成绩表的cno数据删除
	err = DefaultGrade.DeleteGradeInfoByCno(sess, cno)
	if err != nil {
		sess.Rollback()
		return false
	}

	err = sess.Commit()
	if err != nil {
		return false
	}
	return true
}

func (this CourseLogic) CourseExist(cno int) bool {
	c := &model.Course{}
	exist, _ := MasterDB.Where("cno=?", cno).Exist(c)
	return exist
}
