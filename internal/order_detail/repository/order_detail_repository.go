package order_detail

import (
	entity "online-course/internal/order_detail/entity"

	"gorm.io/gorm"
)

type OrderDetailRepository interface {
	Create(entity entity.OrderDetail) (*entity.OrderDetail, error)
}

type OrderDetailRepositoryImpl struct {
	db *gorm.DB
}

// Create implements OrderDetailRepository
func (repository *OrderDetailRepositoryImpl) Create(entity entity.OrderDetail) (*entity.OrderDetail, error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

func NewOrderDetailRepository(db *gorm.DB) OrderDetailRepository {
	return &OrderDetailRepositoryImpl{db}
}
