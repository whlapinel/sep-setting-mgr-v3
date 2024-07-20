package auth

import (
	"log"
	"net/http"
	"sep_setting_mgr/internal/domain/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

const sessionLifeSpan = time.Minute * 60

type jwtCustomClaims struct {
	Email string `json:"email"`
	ID    int    `json:"id"`
	jwt.RegisteredClaims
}

var config = echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(jwtCustomClaims)
	},
	SigningKey: []byte("secret"),
	ErrorHandler: func(c echo.Context, err error) error {
		// Redirect to login page on error
		return c.Redirect(303, "/unauthorized")
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
	cookie.Expires = time.Now().Add(time.Minute * 5)
	log.Println("Setting cookie: ", cookie)
	c.SetCookie(cookie)
	c.Response().Header().Set("Authorization", "Bearer "+t)
}

var AddCookieToHeader = func(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		log.Println("running AddCookieToHeader middleware")
		cookie, err := c.Cookie("token")
		if err != nil {
			return c.Redirect(303, "/unauthorized")
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

func Authorization(userRepo models.UserRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID := c.Get("id").(int)
			log.Println("User ID: ", userID)
			user, err := userRepo.FindByID(userID)
			if err != nil {
				return c.String(500, "Failed to get user. See server logs for details.")
			}
			if user == nil {
				return c.String(401, "Unauthorized")
			}
			if !user.Admin {
				return c.Redirect(303, "/unauthorized")
			}
			return next(c)
		}
	}
}
