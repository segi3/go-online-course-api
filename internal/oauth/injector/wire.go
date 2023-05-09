// go:build wireinject

package oauth

import (
	adminRepository "online-course/internal/admin/repository"
	adminUseCase "online-course/internal/admin/usecase"
	oauthHandler "online-course/internal/oauth/delivery/http"
	oauthRepository "online-course/internal/oauth/repository"
	oauthUseCase "online-course/internal/oauth/usecase"
	userRepository "online-course/internal/user/repository"
	userUseCase "online-course/internal/user/usecase"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *oauthHandler.OauthHandler {
	wire.Build(
		oauthHandler.NewOauthHandler,
		oauthRepository.NewOauthClientRepository,
		oauthRepository.NewOauthAccessTokenRepository,
		oauthRepository.NewOauthRefreshTokenRepository,
		oauthUseCase.NewOauthUseCase,
		userRepository.NewUserRepository,
		userUseCase.NewUserUseCase,
		adminRepository.NewAdminRepository,
		adminUseCase.NewAdminUseCase,
	)

	return &oauthHandler.OauthHandler{}

}

// wire gen internal/oauth/injector/wire.go
