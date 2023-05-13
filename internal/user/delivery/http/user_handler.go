package user

import (
	"net/http"
	"strconv"

	"online-course/internal/middleware"
	dto "online-course/internal/user/dto"
	useCase "online-course/internal/user/usecase"
	"online-course/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	useCase useCase.UserUseCase
}

func NewUserHandler(useCase useCase.UserUseCase) *UserHandler {
	return &UserHandler{useCase}
}

func (handler *UserHandler) Route(r *gin.RouterGroup) {
	userHandler := r.Group("/api/v1")

	userHandler.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		userHandler.GET("/users", handler.FindAll)
		userHandler.GET("/users/:id", handler.FindById)
		userHandler.POST("/users", handler.Create)
		userHandler.PATCH("/users/:id", handler.Update)
		userHandler.DELETE("/users/:id", handler.Delete)
	}
}

func (handler *UserHandler) FindAll(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	data := handler.useCase.FindAll(offset, limit)

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}

func (handler *UserHandler) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := handler.useCase.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response(http.StatusNotFound, "not found", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}

func (handler *UserHandler) Create(ctx *gin.Context) {
	var input dto.UserRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.CreatedBy = &user.ID

	data, err := handler.useCase.Create(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, utils.Response(http.StatusCreated, "created", data))
}

func (handler *UserHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.UserRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.UpdatedBy = &user.ID

	data, err := handler.useCase.Update(id, input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}

func (handler *UserHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := handler.useCase.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response(http.StatusNotFound, "not found", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", "ok"))
}
