package order

import (
	"net/http"
	"strconv"

	"online-course/internal/middleware"
	dto "online-course/internal/order/dto"
	useCase "online-course/internal/order/usecase"
	"online-course/pkg/utils"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	useCase useCase.OrderUseCase
}

func NewOrderHandler(useCase useCase.OrderUseCase) *OrderHandler {
	return &OrderHandler{useCase}
}

func (handler *OrderHandler) Route(r *gin.RouterGroup) {
	orderHandler := r.Group("/api/v1")

	orderHandler.Use(middleware.AuthJwt)
	{
		orderHandler.POST("/orders", handler.Create)
		orderHandler.GET("/orders", handler.FindAllByUserId)
	}
}

func (handler *OrderHandler) Create(ctx *gin.Context) {
	var input dto.OrderRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.UserID = user.ID
	input.Email = user.Email

	data, err := handler.useCase.Create(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}

func (handler *OrderHandler) FindAllByUserId(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Param("offset"))
	limit, _ := strconv.Atoi(ctx.Param("limit"))

	user := utils.GetCurrentUser(ctx)

	data := handler.useCase.FindAllByUserId(offset, limit, int(user.ID))

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))

}
