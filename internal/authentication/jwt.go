package authentication

import (
	"rabc-go/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT() (string, error) {
	JwtConfig := config.GetJWTConfig()

	claims := jwt.MapClaims{
		"exp": time.Now().Add(24 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
		// add more claims later
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JwtConfig.Secret))
}
