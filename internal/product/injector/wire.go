//go:build wireinject
// +build wireinject

package product

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	handler "online-course/internal/product/delivery/http"
	repository "online-course/internal/product/repository"
	usecase "online-course/internal/product/usecase"
	fileUpload "online-course/pkg/fileupload/cloudinary"
)

func InitiliazedService(db *gorm.DB) *handler.ProductHandler {
	wire.Build(
		handler.NewProductHandler,
		usecase.NewProductUseCase,
		repository.NewProductRepository,
		fileUpload.NewFileUpload,
	)

	return &handler.ProductHandler{}
}
