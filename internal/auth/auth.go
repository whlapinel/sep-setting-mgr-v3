package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sep_setting_mgr/internal/domain/models"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/idtoken"
)

const SessionLifeSpan = time.Hour
const cushionTime = time.Minute * 5

var secret = os.Getenv("JWT_SECRET")

type jwtCustomClaims struct {
	First   string `json:"first"`
	Last    string `json:"last"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
	ID      int    `json:"id"`
	jwt.RegisteredClaims
}

func AddCookieToHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("running AddCookieToHeader middleware")
		cookie, err := c.Cookie("token")
		if err != nil {
			reason := "Not signed in"
			return c.Redirect(303, unauthorizedPath(c, reason, 0))
		}
		c.Request().Header.Set("Authorization", "Bearer "+cookie.Value)
		return next(c)
	}
}

var JWTMiddleware = echojwt.WithConfig(config)

var config = echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(jwtCustomClaims)
	},
	SigningKey: []byte(secret),
	ErrorHandler: func(c echo.Context, err error) error {
		// Redirect to login page on error
		log.Println("Error: ", err)
		reason := "error validating token"
		return c.Redirect(303, unauthorizedPath(c, reason, 0))
	},
	SuccessHandler: func(c echo.Context) {
		log.Println("SuccessHandler")
		log.Println("Claims: ", c.Get("user"))
		userToken := c.Get("user").(*jwt.Token)
		claims := userToken.Claims.(*jwtCustomClaims)
		log.Println("Email: ", claims.Email)
		log.Println("ID: ", claims.ID)
		log.Println("Name: ", claims.First, claims.Last)
		log.Println("Picture: ", claims.Picture)
		expiration := claims.ExpiresAt.Time
		log.Println("Expiration: ", expiration)
		if time.Until(expiration) <= cushionTime {
			log.Println("less than a minute left")
			t, err := IssueToken(claims.First, claims.Last, claims.Email, claims.Picture, claims.ID)
			if err != nil {
				log.Println("Failed to issue token: ", err)
			}
			WriteToken(c, t)
			jsonString := fmt.Sprintf("{\"signin\":{\"expiration\":%d}}", time.Now().Add(SessionLifeSpan).UnixMilli())
			c.Response().Header().Set("Hx-Trigger", jsonString)
		} else {
			log.Println("more than a minute left")
		}
	},
}

func GetClaims(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("running GetClaims middleware")
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)
		c.Set("first", claims.First)
		c.Set("last", claims.Last)
		c.Set("email", claims.Email)
		c.Set("picture", claims.Picture)
		c.Set("id", claims.ID)
		return next(c)
	}
}

func Authorization(userRepo models.UserRepository, role models.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log.SetPrefix("Authorization Middleware")
			userID := c.Get("id").(int)
			log.Println("User ID: ", userID)
			log.Println("role: ", role)
			user, err := userRepo.FindByID(userID)
			log.Println("user.Admin: ", user.Admin)
			ok := true
			var reason UnauthReason
			if err != nil {
				ok = false
				reason = ErrorRetrievingUser
			} else if user == nil {
				ok = false
				reason = UserNotFound
			} else if role == "admin" {
				if !user.Admin {
					ok = false
					reason = NoAdminRole
				}
			} else if role == "teacher" {
				if !user.Teacher {
					ok = false
					reason = NoTeacherRole
				}
			}
			if !ok {
				return c.Redirect(303, unauthorizedPath(c, reason.String(), userID))
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

func IssueToken(firstName, lastName, email, picture string, id int) (string, error) {
	claims := jwtCustomClaims{
		First:   firstName,
		Last:    lastName,
		Email:   email,
		Picture: picture,
		ID:      id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(SessionLifeSpan)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
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
	cookie.Expires = time.Now().Add(SessionLifeSpan)
	log.Println("Setting cookie: ", cookie)
	c.SetCookie(cookie)
	c.Response().Header().Set("Authorization", "Bearer "+t)
}

func IsSignedIn(c echo.Context) bool {
	cookie, err := c.Cookie("token")
	if err != nil {
		return false
	}
	return cookie.Value != ""
}

func unauthorizedPath(c echo.Context, reason string, userID int) string {
	log.SetPrefix("UnauthorizedPath: ")
	trimmedURI := strings.Trim(c.Request().RequestURI, "/")
	splitURI := strings.Split(trimmedURI, "/")
	var page string
	if len(splitURI) > 0 {
		page = splitURI[0]
	} else {
		page = trimmedURI
	}
	escapedReason := url.QueryEscape(reason)
	return "/unauthorized" + "/" + page + "/" + escapedReason + "/" + strconv.Itoa(userID)
}

type UnauthReason string

func (r UnauthReason) String() string {
	return string(r)
}

const NoAdminRole UnauthReason = "user does not have admin role"
const NoTeacherRole UnauthReason = "user does not have teacher role"
const UserNotFound UnauthReason = "user not found"
const ErrorRetrievingUser UnauthReason = "error retrieving user"
