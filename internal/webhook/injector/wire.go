//go:build wireinject
// +build wireinject

package webhook

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	cartRepository "online-course/internal/cart/repository"
	cartUseCase "online-course/internal/cart/usecase"
	classRoomRepository "online-course/internal/class_room/repository"
	classRoomUseCase "online-course/internal/class_room/usecase"
	discountRepository "online-course/internal/discount/repository"
	discountUseCase "online-course/internal/discount/usecase"
	orderRepository "online-course/internal/order/repository"
	orderUseCase "online-course/internal/order/usecase"
	orderDetailRepository "online-course/internal/order_detail/repository"
	orderDetailUseCase "online-course/internal/order_detail/usecase"
	paymentUseCase "online-course/internal/payment/usecase"
	productRepository "online-course/internal/product/repository"
	productUseCase "online-course/internal/product/usecase"
	handler "online-course/internal/webhook/delivery/http"
	useCase "online-course/internal/webhook/usecase"
	fileUpload "online-course/pkg/fileupload/cloudinary"
)

func InitializedService(db *gorm.DB) *handler.WebhookHandler {
	wire.Build(
		handler.NewWebHookHandler,
		useCase.NewWebhookUseCase,
		classRoomRepository.NewClassRoomRepository,
		classRoomUseCase.NewClassRoomUseCase,
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
	)

	return &handler.WebhookHandler{}
}
