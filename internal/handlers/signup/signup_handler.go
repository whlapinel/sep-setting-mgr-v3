package signup

import (
	"log"
	"os"
	"sep_setting_mgr/internal/auth"
	common "sep_setting_mgr/internal/handlers/handlerscommon"
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/handlers/views/layouts"
	"sep_setting_mgr/internal/services/signup"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type SignupHandler interface {
	// GET /signup
	SignUpPage(c echo.Context) error

	// POST /hx-signup
	// Signup(c echo.Context) error

	// POST /signup
	GoogleSignup(c echo.Context) error
}

type handler struct {
	service signup.SignupService
}

func NewHandler(svc signup.SignupService) SignupHandler {
	return &handler{service: svc}
}

var router *echo.Echo

func Mount(e *echo.Echo, h SignupHandler) {
	router = e
	e.GET("/signup", h.SignUpPage).Name = string(common.SignupPage)
	// e.POST("/hx-signup", h.Signup).Name = string(common.Signup)
	e.POST("/signup", h.GoogleSignup).Name = string(common.GoogleSignup)
}

// func (h handler) Signup(c echo.Context) error {
// 	email := c.FormValue("email")
// 	name := c.FormValue("name")
// 	created, err := h.service.CreateUser(first, name)
// 	if err != nil {
// 		log.Println("SignUpHandler(): ", err)
// 		c.String(500, "Error creating user.")
// 		return err
// 	}
// 	if !created {
// 		fmt.Println("User not created.")
// 		return echo.ErrInternalServerError
// 	}
// 	return c.String(201, "User created.")
// }

func (h handler) SignUpPage(c echo.Context) error {
	var clientID string = os.Getenv("GOOGLE_CLIENT_ID")
	if util.IsHTMX(c) {
		return util.RenderTempl(views.SignUpPage(router, clientID), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(views.SignUpPage(router, clientID), nil), c, 200)
}

func (h handler) GoogleSignup(c echo.Context) error {
	payload, err := auth.GoogleAuth(c)
	if err != nil {
		log.Println("Failed to authenticate user: ", err)
		return c.String(401, "Failed to authenticate user")
	}
	email := payload.Claims["email"].(string)
	first := payload.Claims["given_name"].(string)
	last := payload.Claims["family_name"].(string)
	picture := payload.Claims["picture"].(string)

	if !isValidEmail(email) {
		log.Println("Invalid email: ", email)
		return c.String(401, "Invalid email")
	}

	// Check if the user is already in the database
	isDuplicate, err := h.service.IsDuplicate(email)
	if err != nil {
		log.Println("Failed to check for duplicate email: ", err)
		return c.String(500, "Failed to check for duplicate email")
	}
	if isDuplicate {
		log.Println("Email already exists: ", email)
		return c.String(409, "Email already exists")
	}

	// create user
	created, err := h.service.CreateUser(first, last, email, picture)
	if err != nil {
		log.Println("Failed to create user: ", err)
		return c.String(500, "Failed to create user")
	}
	if !created {
		log.Println("Failed to create user")
		return c.String(500, "Failed to create user")
	}
	return c.String(200, "Welcome "+first+" "+last+"!")
}

func isValidEmail(email string) bool {
	return verifyCMSDomain(email) || email == "whlapinel@gmail.com"
}

func verifyCMSDomain(email string) bool {
	// Check if the email is from a CMS domain
	cmsDomain := "cms.k12.nc.us"
	return email[len(email)-len(cmsDomain):] == cmsDomain
}
