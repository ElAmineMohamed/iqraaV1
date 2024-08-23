package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/elaminemohamed/iqraa/Type"
	"github.com/elaminemohamed/iqraa/views"
	"github.com/labstack/echo/v4"
)

type UsersResponse struct {
	Users []Type.User `json:"users"`
}

func HandelHome(c echo.Context) error {
	res, _ := http.Get("https://dummyjson.com/user")
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var response UsersResponse
	_ = json.Unmarshal(body, &response)
	hx := c.Request().Header.Get("Hx-Request")
	return Render(c, 200, views.Home(hx, &response.Users))
}

func HandelAbout(c echo.Context) error {
	hx := c.Request().Header.Get("Hx-Request")
	return Render(c, 200, views.About(hx))
}
func HandelPuzzl(c echo.Context) error {
	hx := c.Request().Header.Get("Hx-Request")
	return Render(c, 200, views.Puzzl(hx))
}
func HandelTeacher(c echo.Context) error {
	hx := c.Request().Header.Get("Hx-Request")
	return Render(c, 200, views.Teachre(hx))
}
func HandelCourse(c echo.Context) error {
	hx := c.Request().Header.Get("Hx-Request")
	return Render(c, 200, views.Course(hx))
}
