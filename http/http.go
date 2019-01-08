package http

import (
	"encoding/json"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func Success(ctx echo.Context, data interface{}) error {
	result := map[string]interface{}{
		"code": 1,
		"msg":  "操作成功",
		"data": data,
	}

	b, err := json.Marshal(result)
	if err != nil {
		return err
	}

	return ctx.JSONBlob(http.StatusOK, b)
}

func Fail(ctx echo.Context, code int, msg string) error {
	result := map[string]interface{}{
		"code": code,
		"msg":  msg,
	}

	log.Println("operate fail:", result)
	return ctx.JSON(http.StatusOK, result)
}
