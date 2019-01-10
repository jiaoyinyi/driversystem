package logic

import (
	. "driversystem/db"
	"driversystem/model"
	"errors"
	"github.com/labstack/echo"
)

type GradeLogic struct{}

var DefaultGrade = GradeLogic{}

func (this GradeLogic) CreateGrade(g *model.Grade) error {
	exist := this.GradeExist(g.Sno, g.Cno)
	if exist {
		return errors.New("grade info exist")
	}
	_, err := MasterDB.Insert(g)
	return err
}

func (this GradeLogic) GetGradeInfoBySnoAndCno(sno, cno int) *model.Grade {
	g := &model.Grade{}
	ok, _ := MasterDB.Where("sno=? AND cno=?", sno, cno).Get(g)
	if !ok {
		return nil
	}
	return g
}

func (this GradeLogic) GetGradeInfos(ctx echo.Context) []*model.Grade {
	offset, limit := DefaultPage.GetPage(ctx)

	gs := make([]*model.Grade, 0)
	MasterDB.Limit(limit, offset).Find(&gs)
	return gs
}

func (this GradeLogic) GetGradeInfoCount() int {
	g := &model.Grade{}
	num, _ := MasterDB.Count(g)
	return int(num)
}

func (this GradeLogic) GetGradeInfoBySno(ctx echo.Context, sno int) []*model.Grade {
	offset, limit := DefaultPage.GetPage(ctx)

	gs := make([]*model.Grade, 0)
	MasterDB.Where("sno=?", sno).Limit(limit, offset).Find(&gs)
	return gs
}

func (this GradeLogic) GetGradeInfoCountBySno(sno int) int {
	g := &model.Grade{}
	num, _ := MasterDB.Where("sno=?", sno).Count(g)
	return int(num)
}

func (this GradeLogic) GetGradeInfoByCno(ctx echo.Context, cno int) []*model.Grade {
	offset, limit := DefaultPage.GetPage(ctx)

	gs := make([]*model.Grade, 0)
	MasterDB.Where("cno=?", cno).Limit(limit, offset).Find(&gs)
	return gs
}

func (this GradeLogic) GetGradeInfoCountByCno(cno int) int {
	g := &model.Grade{}
	num, _ := MasterDB.Where("cno=?", cno).Count(g)
	return int(num)
}

func (this GradeLogic) UpdateGradeInfo(g *model.Grade, cols []string) bool {
	num, _ := MasterDB.Where("id=?", g.Id).Cols(cols...).Update(g)
	if num == 0 {
		return false
	}
	return true
}

func (this GradeLogic) DeleteGradeInfo(id int) bool {
	g := &model.Grade{}
	num, _ := MasterDB.Where("id=?", id).Delete(g)
	if num == 0 {
		return false
	}
	return true
}

func (this GradeLogic) GradeExist(sno, cno int) bool {
	g := &model.Grade{}
	exist, _ := MasterDB.Where("sno=? AND cno=?", sno, cno).Exist(g)
	return exist
}

func (this GradeLogic) DealGrade(g *model.Grade) *model.GradeInfo {
	info := &model.GradeInfo{
		Id:       g.Id,
		Sno:      g.Sno,
		Cno:      g.Cno,
		LastTime: FormatTime(g.LastTime),
		Times:    g.Times,
		Grade:    g.Grade,
	}
	return info
}

func (this GradeLogic) DealGrades(gs []*model.Grade) []*model.GradeInfo {
	infos := make([]*model.GradeInfo, len(gs))
	for index, g := range gs {
		infos[index] = this.DealGrade(g)
	}
	return infos
}
