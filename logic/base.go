package logic

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"time"
)

func getSession(ctx echo.Context) *sessions.Session {
	sess, _ := session.Get("driversystem", ctx)
	return sess
}

func NewCookieSession(ctx echo.Context) (*sessions.Session) {
	sess := getSession(ctx)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
	}
	sess.Save(ctx.Request(), ctx.Response())
	return sess
}

func SetSessionValues(ctx echo.Context, pairs map[string]interface{}) {
	sess := getSession(ctx)
	for key, val := range pairs {
		sess.Values[key] = val
	}
	sess.Save(ctx.Request(), ctx.Response())
}

func GetSessionValue(ctx echo.Context, key string) interface{} {
	sess := getSession(ctx)
	return sess.Values[key]
}

func DeleteSessionValue(ctx echo.Context, key string) {
	sess := getSession(ctx)
	delete(sess.Values, key)
	sess.Save(ctx.Request(), ctx.Response())
}

func FormatTime(time time.Time) string {
	formatTime := time.Format("2006-01-02")
	return formatTime
}

func ParseTime(str string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", str)
	return t, err
}
