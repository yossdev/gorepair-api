package middleware

import (
	"fmt"
	"gorepair-rest-api/config"
	"gorepair-rest-api/internal/web"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func UserRestricted(ctx *fiber.Ctx) error {
	jwtToken := strings.Replace(ctx.Get("Authorization"), fmt.Sprintf("%s ", config.Get().JwtTokenType), "", 1)

	token, _ := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		tokenType := t.Claims.(jwt.MapClaims)["token_type"]
		if tokenType != "refresh_token" {
			return nil, fmt.Errorf("unexpected token type: %v", tokenType)
		}
		return []byte(config.Get().AppKey), nil
	})

	role := token.Claims.(jwt.MapClaims)["role"].(string)

	if role != "user" {
		return web.JsonResponse(ctx, http.StatusForbidden, "forbidden", nil)
	}

	return ctx.Next()
}

func WorkshopRestricted(ctx *fiber.Ctx) error {
	jwtToken := strings.Replace(ctx.Get("Authorization"), fmt.Sprintf("%s ", config.Get().JwtTokenType), "", 1)

	token, _ := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		tokenType := t.Claims.(jwt.MapClaims)["token_type"]
		if tokenType != "refresh_token" {
			return nil, fmt.Errorf("unexpected token type: %v", tokenType)
		}
		return []byte(config.Get().AppKey), nil
	})

	role := token.Claims.(jwt.MapClaims)["role"].(string)

	if role != "workshop" {
		return web.JsonResponse(ctx, http.StatusForbidden, "forbidden", nil)
	}

	return ctx.Next()
}