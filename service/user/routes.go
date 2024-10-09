package user

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sysronnie/go-user-auth/service/auth"
	"github.com/sysronnie/go-user-auth/types"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(group *echo.Group) {
	group.POST("/register", h.RegisterUser)
	group.POST("/login", h.HandleLogin)
}


type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Cookie bool `json:"cookie"`
}

func (h *Handler) HandleLogin(c echo.Context) error {
	var req LoginRequest

	// Bind JSON body to the struct
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

	username := req.Username
	password := req.Password
	cookie := req.Cookie


	u, err := h.store.GetUserByUsername(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error retrieving user",
		})
	}

	if !auth.ComparePasswords(u.PASSWORD, []byte(password)) {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid username or password",
		})
	}

	secret := "admin"
	token, err := auth.CreateJWT([]byte(secret), u.USER_ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Could not create token",
		})
	}

	if (cookie) {
		auth.SetJWTTokenCookie(c, token, time.Hour*24)
	} 

	// Set token in response body for mobile clients
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"token":   token,
	})
}


// Input validation example
func (h *Handler) RegisterUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	hashedPassword, err := auth.HashPassword(password)

	if err != nil {
		log.Fatal("wow")
	}

	email := c.FormValue("email")
	User := new(types.User)
	User = &types.User{
		USERNAME: username,
		PASSWORD: hashedPassword,
		EMAIL: email,
	}

	h.store.CreateUser(*User)
	

	return c.JSON(http.StatusOK, map[string]string{
		"username": username,
		"password": password,  // Don't return the password in production
	})
}
