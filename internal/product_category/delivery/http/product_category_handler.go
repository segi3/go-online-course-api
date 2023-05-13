package product_category

import (
	"net/http"
	"strconv"

	"online-course/internal/middleware"
	dto "online-course/internal/product_category/dto"
	usecase "online-course/internal/product_category/usecase"
	"online-course/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ProductCategoryHandler struct {
	usecase usecase.ProductCategoryUseCase
}

func NewProductCategoryHandler(usecase usecase.ProductCategoryUseCase) *ProductCategoryHandler {
	return &ProductCategoryHandler{usecase}
}

func (handler *ProductCategoryHandler) Route(r *gin.RouterGroup) {
	productCategoryRouter := r.Group("/api/v1")

	productCategoryRouter.GET("/product_categories", handler.FindAll)
	productCategoryRouter.GET("/product_categories/:id", handler.FindById)

	productCategoryRouter.Use(middleware.AuthJwt, middleware.AuthAdmin)
	{
		productCategoryRouter.POST("/product_categories", handler.Create)
		productCategoryRouter.PATCH("/product_categories/:id", handler.Update)
		productCategoryRouter.DELETE("/product_categories/:id", handler.Delete)
	}
}

func (handler *ProductCategoryHandler) Create(ctx *gin.Context) {
	var input dto.ProductCategoryRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.CreatedBy = user.ID

	productCategory, err := handler.usecase.Create(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", productCategory))
}

func (handler *ProductCategoryHandler) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.ProductCategoryRequestBody

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response(http.StatusBadRequest, "bad request", err.Error()))
		ctx.Abort()
		return
	}

	user := utils.GetCurrentUser(ctx)

	input.UpdatedBy = user.ID

	// Lakukan update
	productCategory, err := handler.usecase.Update(id, input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response(http.StatusInternalServerError, "internal server error", err.Error()))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", productCategory))
}

func (handler *ProductCategoryHandler) FindAll(ctx *gin.Context) {
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	productCategories := handler.usecase.FindAll(offset, limit)

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", productCategories))
}

func (handler *ProductCategoryHandler) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	productCategory, err := handler.usecase.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response(http.StatusNotFound, "data not found", "data not found"))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", productCategory))
}

func (handler *ProductCategoryHandler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := handler.usecase.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, utils.Response(http.StatusNotFound, "data not found", "data not found"))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, utils.Response(http.StatusOK, "ok", "success"))
}
