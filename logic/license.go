package logic

import (
	. "driversystem/db"
	"driversystem/model"
)

type LicenseLogic struct{}

var DefaultLicense = LicenseLogic{}

func (this LicenseLogic) CreateLicenseInfo(l *model.License, cols []string) error {

}

func (this LicenseLogic) FindOneBySno(sno int) *model.License {

}

func (this LicenseLogic) LicenseExist(sno int) bool {
	l := &model.License{}
	MasterDB.Where("sno=?", sno).Exist(l)
}
