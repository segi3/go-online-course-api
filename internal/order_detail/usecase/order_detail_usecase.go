package order_detail

import (
	"fmt"

	entity "online-course/internal/order_detail/entity"
	repository "online-course/internal/order_detail/repository"
)

type OrderDetailUseCase interface {
	Create(entity entity.OrderDetail)
}

type OrderDetailUseCaseImpl struct {
	repository repository.OrderDetailRepository
}

// Create implements OrderDetailUseCase
func (usecase *OrderDetailUseCaseImpl) Create(entity entity.OrderDetail) {
	_, err := usecase.repository.Create(entity)

	if err != nil {
		fmt.Print(err)
	}
}

func NewOrderDetailUseCase(repository repository.OrderDetailRepository) OrderDetailUseCase {
	return &OrderDetailUseCaseImpl{repository}
}
