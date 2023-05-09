package admin

import (
	handler "online-course/internal/admin/delivery/http"
	repository "online-course/internal/admin/repository"
	usecase "online-course/internal/admin/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *handler.AdminHandler {
	wire.Build(
		repository.NewAdminRepository,
		usecase.NewAdminUseCase,
		handler.NewAdminHandler,
	)

	return &handler.AdminHandler{}
}
