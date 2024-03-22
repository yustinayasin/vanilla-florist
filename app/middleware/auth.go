package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"vanilla-florist/utils"

	"github.com/golang-jwt/jwt/v5"

	users "vanilla-florist/business/user"
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

type GeneratorToken interface {
	GenerateToken(userId int) string
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

func verifyToken(tokenString string, jwtConf ConfigJWT, c http.ResponseWriter) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConf.SecretJWT), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		utils.ReturnErrorResponse(c, http.StatusForbidden, "Token is invalid")
		return nil, nil
	}

	return token, nil
}

// Auth for private routes
func RequireAuth(next http.HandlerFunc, jwtConf ConfigJWT) http.HandlerFunc {
	return func(c http.ResponseWriter, r *http.Request) {
		// Get the token from header
		c.Header().Set("Content-Type", "application/json")
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			c.WriteHeader(http.StatusUnauthorized)
			utils.ReturnErrorResponse(c, http.StatusForbidden, "Unauthorize")
			return
		}

		tokenString = tokenString[len("Bearer "):]

		// Verify token
		token, err := verifyToken(tokenString, jwtConf, c)
		if err != nil {
			c.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(c, "Invalid token")
			return
		}

		// Check the expiry date
		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			http.Error(c, "Failed to extract claims", http.StatusUnauthorized)
			return
		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			http.Error(c, "Internal Server Error", http.StatusUnauthorized)
			return
		}

		// Find the user
		// Access the "subs" claim and convert it to int
		subsClaim, ok := claims["subs"].(float64)

		if !ok {
			fmt.Println("subs claim is not a valid float64")
			return
		}

		// Convert subsClaim to int
		subsInt := int(subsClaim)

		userUseCase := &users.UserUseCase{}

		user, err := userUseCase.FindUser(subsInt)

		if err != nil {
			http.Error(c, "Failed to fetch user data", http.StatusInternalServerError)
			return
		}

		// Attach the user to the request context
		ctx := context.WithValue(r.Context(), "user", user)

		// Pass the updated context to the next handler
		next.ServeHTTP(c, r.WithContext(ctx))
	}
}
