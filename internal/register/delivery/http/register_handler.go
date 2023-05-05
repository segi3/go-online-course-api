package register

import (
	"net/http"
	registerUseCase "online-course/internal/register/usecase"
	userDto "online-course/internal/user/dto"
	utils "online-course/pkg/utils"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	registerUseCase registerUseCase.RegisterUseCase
}

func NewRegisterHandler(registerUseCase registerUseCase.RegisterUseCase) *RegisterHandler {
	return &RegisterHandler{registerUseCase}
}

func (rh *RegisterHandler) Route(r *gin.RouterGroup) {
	r.POST("/api/v1/register", rh.Register)
}

func (rh *RegisterHandler) Register(ctx *gin.Context) {

	// request body validation
	var registerRequestInput userDto.UserRequestBody

	if err := ctx.ShouldBindJSON(&registerRequestInput); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(400, "Bad Request", err.Error()))
		ctx.Abort()
		return
	}

	// register
	err := rh.registerUseCase.Register(registerRequestInput)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(500, "Internal server error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, utils.Response(201, "created", "Registration success"))
}
