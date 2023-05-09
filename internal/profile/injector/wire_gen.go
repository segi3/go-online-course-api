// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package profile

import (
	"gorm.io/gorm"
	"online-course/internal/profile/delivery/http"
	profile2 "online-course/internal/profile/usecase"
	"online-course/internal/user/repository"
	user2 "online-course/internal/user/usecase"
)

// Injectors from wire.go:

func InitializeService(db *gorm.DB) *profile.ProfileHandler {
	userRepository := user.NewUserRepository(db)
	userUseCase := user2.NewUserUseCase(userRepository)
	profileUseCase := profile2.NewProfileUseCase(userUseCase)
	profileHandler := profile.NewProfileHandler(profileUseCase)
	return profileHandler
}
