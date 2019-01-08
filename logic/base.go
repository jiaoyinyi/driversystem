package logic

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
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

//const (
//	TokenSalt = "1sd2HKge.643!SF2dgfdg4*"
//	TokenLen  = 26
//)
//
//func genScrect(uid int) []byte {
//	//时间+TokenSalt+uid
//	strT := strconv.FormatInt(time.Now().Unix(), 10)
//	strUid := strconv.FormatInt(int64(uid), 10)
//
//	buffer := bytes.Buffer{}
//	buffer.WriteString(strT)
//	buffer.WriteString(TokenSalt)
//	buffer.WriteString(strUid)
//
//	return buffer.Bytes()
//}
//
////生成token
//func GenToken(uid int) string {
//	h := md5.New()
//	h.Write(genScrect(uid))
//	body := hex.EncodeToString(h.Sum(nil))
//
//	pt := time.Now().Add(time.Hour).Unix()
//	strPT := strconv.FormatInt(pt, 10)
//
//	buffer := bytes.Buffer{}
//	buffer.WriteString(strPT)
//	buffer.WriteString(body)
//	token := buffer.String()
//	return token
//}
//
//func ValidateToken(token string) bool {
//	if len(token) != TokenLen {
//		return false
//	}
//
//	timeStamp, err := strconv.ParseInt(token[:10], 10, 64)
//	if err != nil {
//		return false
//	}
//
//	expireTime := time.Unix(timeStamp, 0)
//	if time.Now().Before(expireTime) {
//		return true
//	}
//	return false
//}
