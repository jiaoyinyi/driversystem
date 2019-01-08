package main

import (
	_ "driversystem/db"
	"driversystem/route"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"log"
)

func main() {
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(securecookie.GenerateRandomKey(32)))))

	g := e.Group("")
	route.RegisterAllRoutes(g)

	err := e.Start(":10006")
	log.Println(err)
}
