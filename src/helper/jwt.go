package helper

import (
	"fmt"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/nazudis/disbursement/src/config"
	"github.com/spf13/viper"
)

// GenerateJWTToken generates a JWT token with custom map claims
func GenerateJWTToken(claims map[string]interface{}) (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	mapClaims := jwt.MapClaims(claims) // Convert claims to jwt.MapClaims
	token.Claims = mapClaims

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(viper.GetString(config.JWTSecret)))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWTToken parses a JWT token and returns the custom claims
func ParseJWTToken(tokenString string) (map[string]interface{}, error) {
	// Parse the token
	var claims jwt.MapClaims
	parser := jwt.NewParser(jwt.WithoutClaimsValidation())
	token, err := parser.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key
		return []byte(viper.GetString(config.JWTSecret)), nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
