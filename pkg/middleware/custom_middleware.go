package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var (
	ErrorContextNotExist = errors.New("user context not exist")
)

func GenerateAccessToken(email, clientID, secret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["email"] = email
	claims["client_id"] = clientID
	claims["nonce"] = time.Now().Add(time.Hour * 1).Unix() // unixnano to milisecond and then convert to string
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// GenerateRefreshToken menghasilkan token penyegar (refresh token)
func GenerateRefreshToken(secret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["nonce"] = time.Now().Add(time.Hour * 24 * 1).Unix() // Token kedaluwarsa dalam 7 hari
	return token.SignedString([]byte(secret))
}

func JwtMiddleware(secretKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// Get token from header
			authorizationHeader := ctx.Request().Header.Get("Authorization")

			// Check if token exists and has correct format
			if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
				response := map[string]interface{}{
					"message": "Bad request",
					"data":    nil,
					"error":   "Invalid token or missing token",
				}
				return ctx.JSON(http.StatusBadRequest, response)
			}

			// Extract token string
			tokenString := strings.Replace(authorizationHeader, "Bearer ", "", 1)

			// Parse token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || method != jwt.SigningMethodHS256 {
					return nil, fmt.Errorf("invalid signing method")
				}
				return []byte(secretKey), nil
			})

			// Check for token parsing errors
			if err != nil || !token.Valid {
				response := map[string]interface{}{
					"message": "Bad request",
					"data":    nil,
					"error":   "Invalid token",
				}
				return ctx.JSON(http.StatusBadRequest, response)
			}

			// Extract user ID from token claims
			claims := token.Claims.(jwt.MapClaims)
			userID := fmt.Sprintf("%v", claims["jti"])
			sessID := fmt.Sprintf("%v", claims["sessid"])

			// Set data to struct
			UserAuthData := UserAuth{
				UserID: userID,
				SessID: sessID,
			}

			// Set to Context
			ctx.Set("UserAuth", &UserAuthData)

			// Proceed to next middleware or handler
			return next(ctx)
		}
	}
}

type UserAuth struct {
	UserID string `json:"user_id"`
	SessID string `json:"sess_id"`
}

// func GetDataFromJWTContext(ctx echo.Context) (*authModel.UserResponse, error) {
// 	payloadClaims, ok := ctx.Get("payload").(*auth.VerifyTokenResponse)
// 	if !ok || payloadClaims == nil {
// 		return nil, ErrorContextNotExist
// 	}

// 	userResp := authModel.UserResponse{
// 		UserID: payloadClaims.Id,
// 		Email:  payloadClaims.Email,
// 		Role:   &payloadClaims.Roles,
// 		SessID: payloadClaims.SessionId,
// 	}

// 	return &userResp, nil
// }
