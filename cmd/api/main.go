package main

import (
	mysql "online-course/pkg/db/mysql"

	"github.com/gin-gonic/gin"

	admin "online-course/internal/admin/injector"
	cart "online-course/internal/cart/injector"
	classRoom "online-course/internal/class_room/injector"
	dashboard "online-course/internal/dashboard/injector"
	discount "online-course/internal/discount/injector"
	oauth "online-course/internal/oauth/injector"
	order "online-course/internal/order/injector"
	product "online-course/internal/product/injector"
	productCategory "online-course/internal/product_category/injector"
	profile "online-course/internal/profile/injector"
	register "online-course/internal/register/injector"
	user "online-course/internal/user/injector"
	webhook "online-course/internal/webhook/injector"
)

func main() {

	// gin.SetMode(gin.ReleaseMode)
	db := mysql.DB()

	r := gin.Default()

	register.InitializedService(db).Route(&r.RouterGroup)
	oauth.InitializedService(db).Route(&r.RouterGroup)
	profile.InitializedService(db).Route(&r.RouterGroup)
	admin.InitializedService(db).Route(&r.RouterGroup)
	productCategory.InitializedService(db).Route(&r.RouterGroup)
	product.InitiliazedService(db).Route(&r.RouterGroup)
	cart.InitiliazedService(db).Route(&r.RouterGroup)
	discount.InitializedService(db).Route(&r.RouterGroup)
	order.InitializedService(db).Route(&r.RouterGroup)
	webhook.InitializedService(db).Route(&r.RouterGroup)
	classRoom.InitializedService(db).Route(&r.RouterGroup)
	dashboard.InitializedService(db).Route(&r.RouterGroup)
	user.InitializedService(db).Route(&r.RouterGroup)

	r.Run()
}
