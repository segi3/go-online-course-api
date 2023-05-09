package product_category

import (
	handler "online-course/internal/product_category/delivery/http"
	repository "online-course/internal/product_category/repository"
	usecase "online-course/internal/product_category/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.ProductCategoryHandler {
	wire.Build(
		handler.NewProductCategoryHandler,
		repository.NewProductCategoryRepository,
		usecase.NewProductCategoryUseCase,
	)

	return &handler.ProductCategoryHandler{}
}
