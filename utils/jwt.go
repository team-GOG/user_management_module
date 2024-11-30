package utils

import (
    "time"

    "github.com/golang-jwt/jwt/v5"
    "github.com/Web-developing-team/user_management_module/config"
)

// GenerateJWT generates a JWT token for the given user ID and role
func GenerateJWT(userID uint, role string) (string, error) {
    claims := jwt.MapClaims{
      "user_id": userID,
      "role":    role,
      "exp":     time.Now().Add(time.Hour * 72).Unix(), // Token expires in 72 hours
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(config.JWTSecret))
}
