package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sysronnie/go-user-auth/service/auth"
	"github.com/sysronnie/go-user-auth/service/user"
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
	// API versioning
	v1 := e.Group("/api/v1")

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(v1)

	// Routes requiring JWT authentication
	authGroup := v1.Group("/auth")
	authGroup.Use(auth.JWTMiddleware("admin")) // Protect these routes with JWT middleware

	authGroup.GET("/protected", s.ProtectedEndpoint) // Example protected endpoint

	return nil
}

func (s *API) ProtectedEndpoint(c echo.Context) error {
	userID, err := auth.GetUserFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid or expired token",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("Welcome User ID: %d", userID),
	})
}
