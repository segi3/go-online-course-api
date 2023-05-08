package middleware

import (
	"net/http"
	utils "online-course/pkg/utils"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	dto "online-course/internal/oauth/dto"
)

type Header struct {
	Authorization string `header: "authorization" binding:"required"`
}

func AuthJwt(ctx *gin.Context) {
	var input Header

	if err := ctx.ShouldBindHeader(input); err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.Response(http.StatusUnauthorized, "unathorized access", err.Error()))
		ctx.Abort()
		return
	}

	reqToken := input.Authorization
	if len(reqToken) != 2 {
		ctx.JSON(http.StatusUnauthorized, utils.Response(http.StatusUnauthorized, "unathorized access", err.Error()))
		ctx.Abort()
		return
	}

	splitToken := strings.Split(reqToken, "Bearer")
	reqToken = splitToken[1]

	claims := &dto.MapClaimsResponse{}

	token, err := jwt.ParseWithClaims(reqToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}
