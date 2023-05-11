package Cart

import (
	handler "online-course/internal/cart/delivery/http"
	repository "online-course/internal/cart/repository"
	usecase "online-course/internal/cart/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitiliazedService(db *gorm.DB) *handler.CartHandler {
	wire.Build(
		handler.NewCartHandler,
		repository.NewCartRepository,
		usecase.NewCartUseCase,
	)

	return &handler.CartHandler{}
}
