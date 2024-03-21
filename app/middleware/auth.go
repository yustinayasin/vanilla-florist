package middleware

import (
	"net/http"
	"time"
	"vanilla-florist/utils"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	UserId int `json:"id"`
	jwt.RegisteredClaims
}

type ConfigJWT struct {
	SecretJWT       string //ambil dari config json
	ExpiresDuration int
}

// Define JWTConfig struct to hold JWT configuration
type JWTConfig struct {
	Claims                  *JwtCustomClaims
	SigningKey              []byte
	ErrorHandlerWithContext func(error, http.ResponseWriter) error
}

// Define a function to initialize JWTConfig
func (jwtConf *ConfigJWT) Init() JWTConfig {
	return JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		ErrorHandlerWithContext: func(e error, c http.ResponseWriter) error {
			// Here you can handle errors however you want
			utils.ReturnErrorResponse(c, http.StatusForbidden, e.Error())
			return nil
		},
	}
}

func (configJwt ConfigJWT) GenerateToken(userId int) string {
	claims := JwtCustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(int64(configJwt.ExpiresDuration)))),
		},
	}

	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := t.SignedString([]byte(configJwt.SecretJWT))

	return token
}

// Define the function to extract user ID from JWT token
func GetUserId(token *jwt.Token) int {
	if token != nil {
		claims, ok := token.Claims.(*JwtCustomClaims)
		if !ok {
			// Token claims couldn't be parsed properly
			return 0
		}
		return claims.UserId
	}
	return 0
}
