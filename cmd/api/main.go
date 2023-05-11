package main

import (
	mysql "online-course/pkg/db/mysql"

	"github.com/gin-gonic/gin"

	registerHandler "online-course/internal/register/delivery/http"
	registerUseCase "online-course/internal/register/usecase"
	userRepository "online-course/internal/user/repository"
	useUseCase "online-course/internal/user/usecase"
	mail "online-course/pkg/mail/sendgrid"

	admin "online-course/internal/admin/injector"
	cart "online-course/internal/cart/injector"
	oauth "online-course/internal/oauth/injector"
	product "online-course/internal/product/injector"
	productCategory "online-course/internal/product_category/injector"
	profile "online-course/internal/profile/injector"
)

func main() {

	db := mysql.DB()

	r := gin.Default()

	mail := mail.NewMail()
	userRepository := userRepository.NewUserRepository(db)
	userUseCase := useUseCase.NewUserUseCase(userRepository)
	registerUseCase := registerUseCase.NewRegisterUseCase(userUseCase, mail)
	registerHandler.NewRegisterHandler(registerUseCase).Route(&r.RouterGroup)

	oauth.InitializeService(db).Route(&r.RouterGroup)
	profile.InitializeService(db).Route(&r.RouterGroup)
	admin.InitializeService(db).Route(&r.RouterGroup)
	productCategory.InitializeService(db).Route(&r.RouterGroup)
	product.InitiliazeService(db).Route(&r.RouterGroup)
	cart.InitiliazeService(db).Route(&r.RouterGroup)

	r.Run()
}
