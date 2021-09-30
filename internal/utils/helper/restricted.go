package helper

import (
	"fmt"
	"gorepair-rest-api/config"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Restricted(ctx *fiber.Ctx) string {
	jwtToken := strings.Replace(ctx.Get("Authorization"), fmt.Sprintf("%s ", config.Get().JwtTokenType), "", 1)

	token, _ := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		tokenType := t.Claims.(jwt.MapClaims)["token_type"]
		if tokenType != "refresh_token" {
			return nil, fmt.Errorf("unexpected token type: %v", tokenType)
		}
		return []byte(config.Get().AppKey), nil
	})

	role := token.Claims.(jwt.MapClaims)["role"].(string)
	// rawID := token.Claims.(jwt.MapClaims)["id"].(float64)

	return role
}