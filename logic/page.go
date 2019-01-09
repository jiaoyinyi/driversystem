package logic

import (
	"github.com/labstack/echo"
	"strconv"
)

var PerPage = 25

type PageLogic struct{}

var DefaultPage = PageLogic{}

func (PageLogic) GetPage(ctx echo.Context) (int, int) {
	var page int64
	p := ctx.QueryParam("page")
	if p != "" {
		page, err := strconv.ParseInt(p, 10, 64)
		if err != nil {
			page = 0
		}
		if page < 0 {
			page = 0
		}
	}
	offset := int(page) * PerPage
	limit := PerPage
	return offset, limit
}
