//go:build wireinject
// +build wireinject

package profile

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	handler "online-course/internal/profile/delivery/http"
	useCase "online-course/internal/profile/usecase"
	userRepository "online-course/internal/user/repository"
	userUseCase "online-course/internal/user/usecase"
)

func InitializedService(db *gorm.DB) *handler.ProfileHandler {
	wire.Build(
		handler.NewProfileHandler,
		useCase.NewProfileUseCase,
		userRepository.NewUserRepository,
		userUseCase.NewUserUseCase,
	)

	return &handler.ProfileHandler{}
}
