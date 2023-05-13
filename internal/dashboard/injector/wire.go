//go:build wireinject
// +build wireinject

package webhook

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	adminRepository "online-course/internal/admin/repository"
	adminUseCase "online-course/internal/admin/usecase"
	cartRepository "online-course/internal/cart/repository"
	cartUseCase "online-course/internal/cart/usecase"
	handler "online-course/internal/dashboard/delivery/http"
	useCase "online-course/internal/dashboard/usecase"
	discountRepository "online-course/internal/discount/repository"
	discountUseCase "online-course/internal/discount/usecase"
	orderRepository "online-course/internal/order/repository"
	orderUseCase "online-course/internal/order/usecase"
	orderDetailRepository "online-course/internal/order_detail/repository"
	orderDetailUseCase "online-course/internal/order_detail/usecase"
	paymentUseCase "online-course/internal/payment/usecase"
	productRepository "online-course/internal/product/repository"
	productUseCase "online-course/internal/product/usecase"
	userRepository "online-course/internal/user/repository"
	userUseCase "online-course/internal/user/usecase"
	fileUpload "online-course/pkg/fileupload/cloudinary"
)

func InitializedService(db *gorm.DB) *handler.DashboardHandler {
	wire.Build(
		handler.NewDashboardHandler,
		useCase.NewDashboardUseCase,
		orderRepository.NewOrderRepository,
		orderUseCase.NewOrderUseCase,
		cartRepository.NewCartRepository,
		cartUseCase.NewCartUseCase,
		discountRepository.NewDiscountRepository,
		discountUseCase.NewDiscountUseCase,
		orderDetailRepository.NewOrderDetailRepository,
		orderDetailUseCase.NewOrderDetailUseCase,
		paymentUseCase.NewPaymentUseCase,
		productRepository.NewProductRepository,
		productUseCase.NewProductUseCase,
		fileUpload.NewFileUpload,
		adminRepository.NewAdminRepository,
		adminUseCase.NewAdminUseCase,
		userRepository.NewUserRepository,
		userUseCase.NewUserUseCase,
	)

	return &handler.DashboardHandler{}
}
