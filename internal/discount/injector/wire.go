//go:build wireinject
// +build wireinject

package discount

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	handler "online-course/internal/discount/delivery/http"
	repository "online-course/internal/discount/repository"
	useCase "online-course/internal/discount/usecase"
)

func InitializedService(db *gorm.DB) *handler.DiscountHandler {
	wire.Build(
		handler.NewDiscountHandler,
		repository.NewDiscountRepository,
		useCase.NewDiscountUseCase,
	)

	return &handler.DiscountHandler{}
}
