package logic

import (
	. "driversystem/db"
	"driversystem/model"
	"errors"
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
)

type LicenseLogic struct{}

var DefaultLicense = LicenseLogic{}

func (this LicenseLogic) CreateLicenseInfo(l *model.License, cols []string) error {
	exist := this.LicenseExist(l.Sno)
	if exist {
		return errors.New("license exist error")
	}

	_, err := MasterDB.Cols(cols...).Insert(l)
	return err
}

func (this LicenseLogic) FindOne(id int) *model.License {
	l := &model.License{}
	ok, _ := MasterDB.Where("id=?", id).Get(l)
	if !ok {
		return nil
	}
	return l
}

func (this LicenseLogic) FindOneBySno(sno int) *model.License {
	l := &model.License{}
	ok, _ := MasterDB.Where("sno=?", sno).Get(l)
	if !ok {
		return nil
	}
	return l
}

func (this LicenseLogic) FindOneByLno(lno string) *model.License {
	l := &model.License{}
	ok, _ := MasterDB.Where("lno=?", lno).Get(l)
	if !ok {
		return nil
	}
	return l
}

func (this LicenseLogic) GetLicenseInfos(ctx echo.Context) []*model.License {
	offset, limit := DefaultPage.GetPage(ctx)

	ls := make([]*model.License, 0)
	_ = MasterDB.Limit(limit, offset).Find(&ls)
	return ls
}

func (this LicenseLogic) GetLicenseInfoCount() int {
	l := &model.License{}
	count, _ := MasterDB.Count(l)
	return int(count)
}

func (this LicenseLogic) GetLicenseInfoBySname(ctx echo.Context, sname string) []*model.License {
	offset, limit := DefaultPage.GetPage(ctx)

	ls := make([]*model.License, 0)
	_ = MasterDB.Where("sname=?", sname).Limit(limit, offset).Find(&ls)
	return ls
}

func (this LicenseLogic) GetLicenseInfoCountBySname(sname string) int {
	l := &model.License{}
	count, _ := MasterDB.Where("sname=?", sname).Count(l)
	return int(count)
}

func (this LicenseLogic) UpdateLicenseInfo(l *model.License, cols []string) bool {
	num, _ := MasterDB.Where("id=?", l.Id).Cols(cols...).Update(l)
	if num == 0 {
		return false
	}
	return true
}

func (this LicenseLogic) DeleteLicenseInfo(id int) bool {
	l := &model.License{}
	num, _ := MasterDB.Where("id=?", id).Delete(l)
	if num == 0 {
		return false
	}
	return true
}

func (this LicenseLogic) DeleteLicenseInfoBySno(sess *xorm.Session, sno int) error {
	l := &model.License{}
	_, err := sess.Where("sno=?", sno).Delete(l)
	return err
}

func (this LicenseLogic) LicenseExist(sno int) bool {
	l := &model.License{}
	exist, _ := MasterDB.Where("sno=?", sno).Exist(l)
	return exist
}

func (this LicenseLogic) DealLicense(l *model.License) *model.LicenseInfo {
	info := &model.LicenseInfo{
		Id:          l.Id,
		Sno:         l.Sno,
		Sname:       l.Sname,
		Lno:         l.Lno,
		ReceiveTime: FormatTime(l.ReceiveTime),
		ReceiveName: l.ReceiveName,
		LText:       l.LText,
	}
	return info
}

func (this LicenseLogic) DealLicenses(ls []*model.License) []*model.LicenseInfo {
	infos := make([]*model.LicenseInfo, len(ls))
	for index, l := range ls {
		infos[index] = this.DealLicense(l)
	}
	return infos
}
