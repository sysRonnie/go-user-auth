package template

import (
	"context"
	"database/sql"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/sysronnie/go-user-auth/types"
	"github.com/sysronnie/go-user-auth/view/page"
)

type TemplateHandler struct {
	DB *sql.DB
}


func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}


func (h TemplateHandler) ShowLandingPage(c echo.Context) error {
	user := types.User{
		USER_ID:   1,
		USERNAME:  "sysRonnie",
	}
	// // Set the user in the context
	// ctx := context.WithValue(c.Request().Context(), "user", user)
	// Then we render our user template, passing the updated context
	return render(c, page.Landing(user))
}


func (h TemplateHandler) ShowHomePage(c echo.Context) error {
	return render(c, page.Home())
}

func (h TemplateHandler) ShowSandboxPage(c echo.Context) error {
	return render(c, page.Home())
}