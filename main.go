package main

import (
	"driversystem/controller"
	_ "driversystem/db"
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

	new(controller.UserController).RegisterRoute(g)

	err := e.Start(":10006")
	log.Println(err)
}
