package oauth

import (
	"fmt"
	"net/http"

	dto "online-course/internal/oauth/dto"
	usecase "online-course/internal/oauth/usecase"
	"online-course/pkg/utils"

	"github.com/gin-gonic/gin"
)

type OauthHandler struct {
	usecase usecase.OauthUseCase
}

func NewOauthHandler(usecase usecase.OauthUseCase) *OauthHandler {
	return &OauthHandler{usecase}
}

func (handler *OauthHandler) Route(r *gin.RouterGroup) {
	oauthRouter := r.Group("/api/v1")

	oauthRouter.POST("/oauth", handler.Login)
}

func (handler *OauthHandler) Login(ctx *gin.Context) {
	var input dto.LoginRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	// Memanggil usecase dari login
	data, err := handler.usecase.Login(input)

	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}
