//go:build wireinject
// +build wireinject

package order

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	cartRepository "online-course/internal/cart/repository"
	cartUseCase "online-course/internal/cart/usecase"
	discountRepository "online-course/internal/discount/repository"
	discountUseCase "online-course/internal/discount/usecase"
	handler "online-course/internal/order/delivery/http"
	repository "online-course/internal/order/repository"
	useCase "online-course/internal/order/usecase"
	orderDetailRepository "online-course/internal/order_detail/repository"
	orderDetailUseCase "online-course/internal/order_detail/usecase"
	paymentUseCase "online-course/internal/payment/usecase"
	productRepository "online-course/internal/product/repository"
	productUseCase "online-course/internal/product/usecase"
	fileUpload "online-course/pkg/fileupload/cloudinary"
)

func InitializedService(db *gorm.DB) *handler.OrderHandler {
	wire.Build(
		cartRepository.NewCartRepository,
		cartUseCase.NewCartUseCase,
		discountRepository.NewDiscountRepository,
		discountUseCase.NewDiscountUseCase,
		handler.NewOrderHandler,
		repository.NewOrderRepository,
		useCase.NewOrderUseCase,
		orderDetailRepository.NewOrderDetailRepository,
		orderDetailUseCase.NewOrderDetailUseCase,
		paymentUseCase.NewPaymentUseCase,
		productRepository.NewProductRepository,
		productUseCase.NewProductUseCase,
		fileUpload.NewFileUpload,
	)

	return &handler.OrderHandler{}
}
