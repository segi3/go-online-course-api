package product

import (
	"net/http"
	"strconv"

	"online-course/internal/middleware"
	dto "online-course/internal/product/dto"
	usecase "online-course/internal/product/usecase"
	"online-course/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	usecase usecase.ProductUseCase
}

func NewProductHandler(usecase usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{usecase}
}

func (handler *ProductHandler) Route(r *gin.RouterGroup) {
	productRoute := r.Group("/api/v1")

	productRoute.GET("/products", handler.FindAll)
	productRoute.GET("/products/:id", handler.FindById)

	productRoute.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		productRoute.POST("/products", handler.Create)
		productRoute.PATCH("/products/:id", handler.Update)
		productRoute.DELETE("/products/:id", handler.Delete)
	}
}

func (handler *ProductHandler) FindAll(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	data := handler.usecase.FindAll(offset, limit)

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}

func (handler *ProductHandler) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := handler.usecase.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response(http.StatusFound, "not found", "not found"))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}

func (handler *ProductHandler) Create(ctx *gin.Context) {
	var input dto.ProductRequestBody

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.CreatedBy = user.ID

	data, err := handler.usecase.Create(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, utils.Response(http.StatusCreated, "created", data))
}

func (handler *ProductHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.ProductRequestBody

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.UpdatedBy = user.ID

	data, err := handler.usecase.Update(id, input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", data))
}

func (handler *ProductHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := handler.usecase.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response(http.StatusFound, "not found", "not found"))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", "success"))
}
