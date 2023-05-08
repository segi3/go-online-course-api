package profile

import (
	"fmt"
	"net/http"
	useCase "online-course/internal/profile/usecase"
	utils "online-course/pkg/utils"
	middleware "online-course/internal/middleware"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	useCase useCase.ProfileUseCase
}

func NewProfileHandler(useCase useCase.ProfileUseCase) *ProfileHandler {
	return &ProfileHandler{useCase}
}

func (handler *ProfileHandler) Route (r *gin.RouterGroup) {
	authorized := r.Group("/api/v1")
	authorized.Use(middleware.AuthJwt){
		authorized.GET("/profiles", handler.GetProfile)
	}
}

func (handler *ProfileHandler) GetProfile(ctx *gin.Context) {
	user := utils.GetCurrentUser(ctx)

	// get profile
	profile, err := handler.useCase.GetProfile(int(user.ID))
	if err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "Bad request", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", profile))
}
