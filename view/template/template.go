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
type userContextKey string

const UserContextKey = userContextKey("user")


func (h TemplateHandler) ShowLandingPage(c echo.Context) error {
	user := types.User{
		USER_ID:   1,
		USERNAME:  "sysRonnie",
	}

	// Set the user in the context
	ctx := context.WithValue(c.Request().Context(), "user", user)

	// Then we render our user template, passing the updated context
	return render(c, page.Landing(user), ctx)
}

func render(c echo.Context, component templ.Component, ctx context.Context) error {
	return component.Render(ctx, c.Response())
}

