package routers

import (
	"github.com/elaminemohamed/iqraa/handlers"
	"github.com/labstack/echo/v4"
)

func Start() {
	e := echo.New()
	e.GET("/", handlers.HandelHome)
	e.GET("/about", handlers.HandelAbout)
	e.GET("/puzzl", handlers.HandelPuzzl)
	e.GET("/teacher", handlers.HandelTeacher)
	e.GET("/course", handlers.HandelCourse)
	e.Static("/static", "static")
	e.Start(":3000")
}
