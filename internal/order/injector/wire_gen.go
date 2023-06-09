// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package order

import (
	"gorm.io/gorm"
	"online-course/internal/cart/repository"
	cart2 "online-course/internal/cart/usecase"
	"online-course/internal/discount/repository"
	discount2 "online-course/internal/discount/usecase"
	"online-course/internal/order/delivery/http"
	order2 "online-course/internal/order/repository"
	order3 "online-course/internal/order/usecase"
	"online-course/internal/order_detail/repository"
	order_detail2 "online-course/internal/order_detail/usecase"
	"online-course/internal/payment/usecase"
	"online-course/internal/product/repository"
	product2 "online-course/internal/product/usecase"
	"online-course/pkg/fileupload/cloudinary"
)

// Injectors from wire.go:

func InitializedService(db *gorm.DB) *order.OrderHandler {
	orderRepository := order2.NewOrderRepository(db)
	cartRepository := cart.NewCartRepository(db)
	cartUseCase := cart2.NewCartUseCase(cartRepository)
	discountRepository := discount.NewDiscountRepository(db)
	discountUseCase := discount2.NewDiscountUseCase(discountRepository)
	productRepository := product.NewProductRepository(db)
	fileUpload := fileupload.NewFileUpload()
	productUseCase := product2.NewProductUseCase(productRepository, fileUpload)
	orderDetailRepository := order_detail.NewOrderDetailRepository(db)
	orderDetailUseCase := order_detail2.NewOrderDetailUseCase(orderDetailRepository)
	paymentUseCase := payment.NewPaymentUseCase()
	orderUseCase := order3.NewOrderUseCase(orderRepository, cartUseCase, discountUseCase, productUseCase, orderDetailUseCase, paymentUseCase)
	orderHandler := order.NewOrderHandler(orderUseCase)
	return orderHandler
}
