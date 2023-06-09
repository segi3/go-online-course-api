// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package oauth

import (
	"gorm.io/gorm"
	"online-course/internal/admin/repository"
	admin2 "online-course/internal/admin/usecase"
	"online-course/internal/oauth/delivery/http"
	oauth2 "online-course/internal/oauth/repository"
	oauth3 "online-course/internal/oauth/usecase"
	"online-course/internal/user/repository"
	user2 "online-course/internal/user/usecase"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *oauth.OauthHandler {
	oauthClientRepository := oauth2.NewOauthClientRepository(db)
	oauthAccessTokenRepository := oauth2.NewOauthAccessTokenRepository(db)
	oauthRefreshTokenRepository := oauth2.NewOauthRefreshTokenRepository(db)
	userRepository := user.NewUserRepository(db)
	userUseCase := user2.NewUserUseCase(userRepository)
	adminRepository := admin.NewAdminRepository(db)
	adminUseCase := admin2.NewAdminUseCase(adminRepository)
	oauthUseCase := oauth3.NewOauthUseCase(oauthClientRepository, oauthAccessTokenRepository, oauthRefreshTokenRepository, userUseCase, adminUseCase)
	oauthHandler := oauth.NewOauthHandler(oauthUseCase)
	return oauthHandler
}
