package signin

import (
	"context"
	"log"
	"net/http"
	"os"
	"sep_setting_mgr/internal/auth"
	"sep_setting_mgr/internal/handlers/common"
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/handlers/views/layouts"
	"sep_setting_mgr/internal/services/signin"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
	"google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type SigninHandler interface {
	// GET /signin
	SignInHandler(e echo.Context) error

	// POST /signin
	GoogleSignin(e echo.Context) error

	// POST /hx-signin
	HxSignin(e echo.Context) error
}

type handler struct {
	service signin.SigninService
}

func NewHandler(svc signin.SigninService) SigninHandler {
	return &handler{service: svc}
}

func Mount(e *echo.Echo, h SigninHandler) {
	e.GET("/signin", h.SignInHandler).Name = string(common.SigninPage)
	e.POST("/signin", h.GoogleSignin).Name = string(common.GoogleSignin)
	e.POST("/hx-signin", h.HxSignin).Name = string(common.SigninPostRoute)
}

func (h handler) SignInHandler(c echo.Context) error {
	isSignedIn := auth.IsSignedIn(c)
	if util.IsHTMX(c) {
		return util.RenderTempl(views.SignInPage(isSignedIn), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(views.SignInPage(isSignedIn)), c, 200)
}

func (h handler) GoogleSignin(c echo.Context) error {
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
	log.Println("User ID: ", tokenInfo.UserId)
	log.Println("Email: ", tokenInfo.Email)
	log.Println("Expires in: ", tokenInfo.ExpiresIn)
	log.Println("Audience: ", tokenInfo.Audience)
	log.Println("StatusCode: ", tokenInfo.HTTPStatusCode)
	log.Println("Header:", tokenInfo.Header)
	log.Println("IssuedTo:", tokenInfo.IssuedTo)
	log.Println("NullFields", tokenInfo.NullFields)
	log.Println("Scope:", tokenInfo.Scope)
	log.Println("ServerResponse:", tokenInfo.ServerResponse)
	log.Println("VerifiedEmail:", tokenInfo.VerifiedEmail)
	// ... Use the user data as needed ...
	// ... Create a session, set cookies, or respond with success ...
	return c.String(200, "GoogleSignin()")
}

func (h handler) HxSignin(c echo.Context) error {
	if !(util.IsHTMX(c)) {
		c.Redirect(303, "/signin")
	}
	email := c.FormValue("email")
	password := c.FormValue("password")
	isAuth, err := h.service.VerifyCredentials(email, password)
	if !isAuth || err != nil {
		return c.String(401, "Invalid credentials")
	}
	id := h.service.GetUserID(email)
	t, err := auth.IssueToken(email, id)
	if err != nil {
		return c.String(500, "Failed to issue token")
	}
	auth.WriteToken(c, t)
	return c.Redirect(303, "/dashboard")
}
