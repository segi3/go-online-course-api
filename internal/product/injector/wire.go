package product

import (
	handler "online-course/internal/product/delivery/http"
	repository "online-course/internal/product/repository"
	usecase "online-course/internal/product/usecase"
	cloudinaryUtils "online-course/pkg/file/cloudinary"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitiliazedService(db *gorm.DB) *handler.ProductHandler {
	wire.Build(
		handler.NewProductHandler,
		usecase.NewProductUseCase,
		repository.NewProductRepository,
		cloudinaryUtils.NewFile,
	)

	return &handler.ProductHandler{}
}
