package signup

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sep_setting_mgr/internal/handlers/common"
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/handlers/views/layouts"
	"sep_setting_mgr/internal/services/signup"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type SignupHandler interface {
	// GET /signup
	SignUpPage(c echo.Context) error

	// POST /hx-signup
	Signup(c echo.Context) error

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
	e.POST("/hx-signup", h.Signup).Name = string(common.Signup)
	e.POST("/signup", h.GoogleSignup).Name = string(common.GoogleSignup)
}

func (h handler) Signup(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	created, err := h.service.CreateUser(email, password)
	if err != nil {
		log.Println("SignUpHandler(): ", err)
		c.String(500, "Error creating user.")
		return err
	}
	if !created {
		fmt.Println("User not created.")
		return echo.ErrInternalServerError
	}
	return c.String(201, "User created.")
}

func (h handler) SignUpPage(c echo.Context) error {
	var clientID string = os.Getenv("GOOGLE_CLIENT_ID")
	if util.IsHTMX(c) {
		return util.RenderTempl(views.SignUpPage(router, clientID), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(views.SignUpPage(router, clientID)), c, 200)
}

func (h handler) GoogleSignup(c echo.Context) error {
	log.Println("GoogleSignup()")
	log.Println(c.Request().Header)
	log.Println(c.Request().Body)
	log.Println(c.Cookie("g_csrf_token"))
	log.Println(c.Cookie("g_state"))
	// ... Get the ID token from the request body or query parameters ...
	idToken := c.FormValue("credential") // Assuming it's sent as "credential" in a form
	log.Println("idToken: ", idToken)

	// Create an OAuth2 service object
	oauth2Service, err := oauth2.NewService(context.Background(), option.WithoutAuthentication())
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to create OAuth2 service")

	}

	// Call the tokeninfo endpoint to verify the token
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return c.String(500, "Failed to verify ID token")
	}
	log.Println(tokenInfo)

	// Extract claims (user information) from the token
	// Verify the "aud" (audience) claim matches your Google client ID
	if tokenInfo.Audience != os.Getenv("GOOGLE_CLIENT_ID") {
		return c.String(401, "Invalid audience")
	}
	// if claims.Aud != os.Getenv("GOOGLE_CLIENT_ID") {
	// 	return c.String(401, "Invalid audience")
	// }

	fmt.Println("User ID: ", tokenInfo.UserId)
	fmt.Println("Email: ", tokenInfo.Email)
	fmt.Println("Name: ", tokenInfo.ExpiresIn)
	fmt.Println("Audience: ", tokenInfo.Audience)
	fmt.Println("StatusCode: ", tokenInfo.HTTPStatusCode)
	fmt.Println("Header:", tokenInfo.Header)
	fmt.Println("IssuedTo:", tokenInfo.IssuedTo)
	fmt.Println("NullFields", tokenInfo.NullFields)
	fmt.Println("Scope:", tokenInfo.Scope)
	fmt.Println("ServerResponse:", tokenInfo.ServerResponse)
	fmt.Println("VerifiedEmail:", tokenInfo.VerifiedEmail)
	// ... Use the user data as needed ...

	// ... Create a session, set cookies, or respond with success ...
	return c.String(200, "GoogleSignup()")
}
