package api

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type API struct {
	db *sql.DB
}

func NewAPIServer(db *sql.DB) *API {
	return &API{
		db: db,
	}
}

func (s *API) APIService(e *echo.Echo) error {
	// Define a group for /api/v1
	v1 := e.Group("/api/v1")

	// Set up a test route
	v1.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "API is working!",
		})
	})

	return nil
}

