package main

import (
	mysql "online-course/pkg/db/mysql"

	"github.com/gin-gonic/gin"

	registerHandler "online-course/internal/register/delivery/http"
	registerUseCase "online-course/internal/register/usecase"
	userRepository "online-course/internal/user/repository"
	useUseCase "online-course/internal/user/usecase"
	mail "online-course/pkg/mail/sendgrid"
)

func main() {

	db := mysql.DB()

	r := gin.Default()

	mail := mail.NewMail()
	userRepository := userRepository.NewUserRepository(db)
	userUseCase := useUseCase.NewUserUseCase(userRepository)
	registerUseCase := registerUseCase.NewRegisterUseCase(userUseCase, mail)
	registerHandler.NewRegisterHandler(registerUseCase).Route(&r.RouterGroup)

	r.Run()
}
