package logic

import (
	. "driversystem/db"
	"driversystem/model"
	"errors"
	"github.com/labstack/echo"
	"log"
)

type HealthLogic struct{}

var DefaultHealth = HealthLogic{}

//插入学员体检信息
func (this HealthLogic) CreateHealthInfo(h *model.Health, cols []string) error {
	log.Println(h)
	exist := this.HealthInfoExist(h.Sno)
	if exist {
		log.Println("exist", h)
		return errors.New("health info exist")
	}
	log.Println("not exist", h)
	_, err := MasterDB.Cols(cols...).Insert(h)
	log.Println(err)
	return err
}

//获取所有学员体检信息 分页
func (this HealthLogic) GetHealthInfos(ctx echo.Context) []*model.Health {
	offset, limit := DefaultPage.GetPage(ctx)

	healths := make([]*model.Health, 0)
	_ = MasterDB.Limit(limit, offset).Find(&healths)
	return healths
}

//获取所有学员体检信息数
func (this HealthLogic) GetHealthInfoCount() int {
	health := &model.Health{}
	count, _ := MasterDB.Count(health)
	return int(count)
}

//根据学员名字获取学员体检信息 分页
func (this HealthLogic) GetHealthInfoBySname(ctx echo.Context, sname string) []*model.Health {
	offset, limit := DefaultPage.GetPage(ctx)

	healths := make([]*model.Health, 0)
	_ = MasterDB.Where("sname=?", sname).Limit(limit, offset).Find(&healths)
	return healths
}

//根据学员名字获取学员体检信息数
func (this HealthLogic) GetHealthInfoCountBySname(sname string) int {
	health := &model.Health{}
	count, _ := MasterDB.Where("sname=?", sname).Count(health)
	return int(count)
}

//根据辩色获取学员体检信息 分页
func (this HealthLogic) GetHealthInfoByDifferentiate(ctx echo.Context, diff string) []*model.Health {
	offset, limit := DefaultPage.GetPage(ctx)

	healths := make([]*model.Health, 0)
	_ = MasterDB.Where("differentiate=?", diff).Limit(limit, offset).Find(&healths)
	return healths
}

//根据辩色获取学员体检信息数
func (this HealthLogic) GetHealthInfoCountByDifferentiate(diff string) int {
	health := &model.Health{}
	count, _ := MasterDB.Where("differentiate=?", diff).Count(health)
	return int(count)
}

func (this HealthLogic) FindOne(id int) *model.Health {
	h := &model.Health{}
	ok, _ := MasterDB.Where("id=?", id).Get(h)
	if !ok {
		return nil
	}
	return h
}

func (this HealthLogic) FindOneBySno(sno int) *model.Health {
	h := &model.Health{}
	ok, _ := MasterDB.Where("sno=?", sno).Get(h)
	if !ok {
		return nil
	}
	return h
}

//删除学员信息
func (this HealthLogic) DeleteHealthInfo(id int) bool {
	h := &model.Health{}
	num, _ := MasterDB.Where("id=?", id).Delete(h)
	if num == 0 {
		return false
	}
	return true
}

func (this HealthLogic) UpdateHealthInfo(h *model.Health, cols []string) bool {
	num, _ := MasterDB.Where("id=?", h.Id).Cols(cols...).Update(h)
	if num == 0 {
		return false
	}
	return true
}

func (this HealthLogic) HealthInfoExist(sno int) bool {
	h := &model.Health{}
	exist, _ := MasterDB.Where("sno=?", sno).Exist(h)
	return exist
}

func (this HealthLogic) CheckDifferentiate(d string) bool {
	_, ok := model.DifferentiateMap[d]
	return ok
}

func (this HealthLogic) CheckLeftEar(le string) bool {
	_, ok := model.LeftEarMap[le]
	return ok
}

func (this HealthLogic) CheckRightEar(re string) bool {
	_, ok := model.RightEarMap[re]
	return ok
}

func (this HealthLogic) CheckLegs(l string) bool {
	_, ok := model.LegsMap[l]
	return ok
}

func (this HealthLogic) CheckPressure(p string) bool {
	_, ok := model.PressureMap[p]
	return ok
}
