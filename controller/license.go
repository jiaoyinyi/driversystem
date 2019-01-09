package controller

import (
	"driversystem/middleware"
	"github.com/labstack/echo"
)

type LicenseController struct{}

func (this LicenseController) RegisterRoute(g *echo.Group) {
	licenseG := g.Group("/license", middleware.NeedLogin)

	licenseG.GET("/search", this.SearchLicenseInfo)
	licenseG.POST("/add", this.AddLicenseInfo)
	licenseG.POST("/update", this.UpdateLicenseInfo)
	licenseG.POST("/delete", this.DeleteLicenseInfo)
}

func (this LicenseController) SearchLicenseInfo(cxt echo.Context) error {
	return nil
}

func (this LicenseController) AddLicenseInfo(cxt echo.Context) error {
	return nil
}

func (this LicenseController) UpdateLicenseInfo(cxt echo.Context) error {
	return nil
}

func (this LicenseController) DeleteLicenseInfo(cxt echo.Context) error {
	return nil
}
