package middlewares

import (
    "net/http"

    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v5"
    "github.com/team-GOG/user_management_module/config"
)

// AuthenticateJWT middleware validates the JWT token and sets user info in context
func AuthenticateJWT(c *fiber.Ctx) error {
    authHeader := c.Get("Authorization")
    if authHeader == "" || len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
      return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid token"})
    }

    tokenString := authHeader[7:]
    token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
      if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fiber.NewError(http.StatusUnauthorized, "Invalid token")
      }
      return []byte(config.JWTSecret), nil
    })

    if err != nil || !token.Valid {
      return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
      return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    // Extract user info from token claims
    c.Locals("user_id", uint(claims["user_id"].(float64)))
    c.Locals("role", claims["role"].(string))
    return c.Next()
}
