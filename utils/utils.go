package utils

import (
	"encoding/json"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)



func render(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}


