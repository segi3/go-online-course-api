//go:build wireinject
// +build wireinject

package register

// mail := mail.NewMail()
// 	userRepository := userRepository.NewUserRepository(db)
// 	userUseCase := userUseCase.NewUserUseCase(userRepository)
// 	registerUseCase := registerUseCase.NewRegisterUseCase(userUseCase, mail)
// 	registerHandler.NewRegisterHandler(registerUseCase).Route(&r.RouterGroup)

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	handler "online-course/internal/register/delivery/http"
	useCase "online-course/internal/register/usecase"
	userRepository "online-course/internal/user/repository"
	userUseCase "online-course/internal/user/usecase"
	mail "online-course/pkg/mail/sendgrid"
)

func InitializedService(db *gorm.DB) *handler.RegisterHandler {
	wire.Build(
		handler.NewRegisterHandler,
		useCase.NewRegisterUseCase,
		userRepository.NewUserRepository,
		userUseCase.NewUserUseCase,
		mail.NewMail,
	)

	return &handler.RegisterHandler{}
}
