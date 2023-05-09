package middleware

import (
	"net/http"

	"online-course/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthAdmin(ctx *gin.Context) {
	admin := utils.GetCurrentUser(ctx)

	if !admin.IsAdmin {
		ctx.JSON(http.StatusUnauthorized, utils.Response(http.StatusUnauthorized, "Unauthorized", "Unauthorized"))
		ctx.Abort()
		return
	}

	ctx.Next()
}
