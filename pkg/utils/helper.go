package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"regexp"
)

const DefaultLanguage = "en"

func GetLanguage(ctx *fiber.Ctx) string {
	return ctx.Get("Accept-Language", DefaultLanguage)
}

func EmailRegex(email string) bool {
	regexpEmail := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
	return regexpEmail.MatchString(email)
}

func GenerateJWT(claims jwt.Claims, method jwt.SigningMethod, jwtSecret string) (string, error) {
	return jwt.NewWithClaims(method, claims).SignedString([]byte(jwtSecret))
}
