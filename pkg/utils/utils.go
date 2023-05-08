package utils

import (
	"math/rand"
	oauthDto "online-course/internal/oauth/dto"

	"github.com/gin-gonic/gin"
)

func RandomString(length int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func GetCurrentUser(ctx *gin.Context) *oauthDto.MapClaimsResponse {
	user, _ := ctx.Get("user")
	return user.(*oauthDto.MapClaimsResponse)
}
