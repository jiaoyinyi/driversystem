package middleware

import (
	. "driversystem/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"log"
)

func NeedLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		cookie, err := ctx.Cookie("driversystem")
		if err != nil {
			return Fail(ctx, 0, "请登录")
		}
		log.Println("cookie:", cookie)

		sess, err := session.Get("driversystem", ctx)
		if err != nil {
			return Fail(ctx, 0, "请登录")
		}

		if sess.IsNew {
			return Fail(ctx, 0, "请登录")
		}

		return next(ctx)
	}
}
