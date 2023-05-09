package utils

import (
	"math/rand"
	oauthDto "online-course/internal/oauth/dto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func Paginate(offset int, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := offset

		// if less than 0, set to 1
		if page <= 0 {
			page = 1
		}

		pageSize := limit

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * limit
		return db.Offset(offset).Limit(pageSize)

	}
}
