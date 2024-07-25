package auth

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
	"sep_setting_mgr/internal/domain/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/idtoken"
)

const sessionLifeSpan = time.Minute * 60

type jwtCustomClaims struct {
	Email string `json:"email"`
	ID    int    `json:"id"`
	jwt.RegisteredClaims
}

type GoogleClaims struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FirstName     string `json:"given_name"`
	LastName      string `json:"family_name"`
	jwt.RegisteredClaims
}

var config = echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(jwtCustomClaims)
	},
	SigningKey: []byte("secret"),
	ErrorHandler: func(c echo.Context, err error) error {
		// Redirect to login page on error
		log.Println("Error: ", err)
		reason := "error validating token"
		escapedReason := url.QueryEscape(reason)
		return c.Redirect(303, "/unauthorized"+c.Request().RequestURI+"/"+escapedReason)
	},
}

func IssueToken(email string, id int) (string, error) {
	claims := jwtCustomClaims{
		Email: email,
		ID:    id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(sessionLifeSpan)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

func WriteToken(c echo.Context, t string) {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(sessionLifeSpan)
	log.Println("Setting cookie: ", cookie)
	c.SetCookie(cookie)
	c.Response().Header().Set("Authorization", "Bearer "+t)
}

var AddCookieToHeader = func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("running AddCookieToHeader middleware")
		cookie, err := c.Cookie("token")
		if err != nil {
			reason := "cookie not found"
			urlEncodedReason := url.QueryEscape(reason)
			return c.Redirect(303, "/unauthorized"+c.Request().RequestURI+"/"+urlEncodedReason)
		}
		c.Request().Header.Set("Authorization", "Bearer "+cookie.Value)
		return next(c)
	}
}

var GetClaims = func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("running GetClaims middleware")
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)
		c.Set("email", claims.Email)
		c.Set("id", claims.ID)
		return next(c)
	}
}
var JWTMiddleware = echojwt.WithConfig(config)

func IsSignedIn(c echo.Context) bool {
	cookie, err := c.Cookie("token")
	if err != nil {
		return false
	}
	return cookie.Value != ""
}

func unauthorizedPath(c echo.Context, reason string, userID int) string {
	escapedReason := url.QueryEscape(reason)
	return "/unauthorized" + c.Request().RequestURI + "/" + escapedReason + "/" + strconv.Itoa(userID)
}

func Authorization(userRepo models.UserRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log.SetPrefix("Authorization Middleware")
			userID := c.Get("id").(int)
			log.Println("User ID: ", userID)
			user, err := userRepo.FindByID(userID)
			if err != nil {
				reason := "error retrieving user"
				return c.Redirect(303, unauthorizedPath(c, reason, userID))
			}
			if user == nil {
				reason := "user not found"
				return c.Redirect(303, unauthorizedPath(c, reason, userID))
			}
			if !user.Admin {
				reason := "not admin"
				return c.Redirect(303, unauthorizedPath(c, reason, userID))
			}
			return next(c)
		}
	}
}

func GoogleAuth(c echo.Context) (*idtoken.Payload, error) {
	token, err := c.Cookie("g_csrf_token")
	if err != nil {
		return nil, errors.New("token not found")
	}
	bodyToken := c.FormValue("g_csrf_token")
	if token.Value != bodyToken {
		return nil, errors.New("token mismatch")
	}
	ctx := context.Background()
	validator, err := idtoken.NewValidator(ctx)
	if err != nil {
		log.Println("Failed to create ID token validator: ", err)
		return nil, errors.New("failed to create ID token validator")
	}
	credential := c.FormValue("credential")
	payload, err := validator.Validate(ctx, credential, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		log.Println("Failed to validate ID token: ", err)
		return nil, errors.New("failed to validate ID token")
	}
	log.Println("Payload: ", payload)
	return payload, nil
}
