package cart

import (
	entity "online-course/internal/cart/entity"
	"online-course/pkg/utils"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindByUserId(userId int, offset int, limit int) []entity.Cart
	FindById(id int) (*entity.Cart, error)
	Create(entity entity.Cart) (*entity.Cart, error)
	Delete(entity entity.Cart) error
	DeleteByUserId(userId int) error
}

type CartRepositoryImpl struct {
	db *gorm.DB
}

// DeleteByUserId implements CartRepository
func (repository *CartRepositoryImpl) DeleteByUserId(userId int) error {
	var cart entity.Cart

	if err := repository.db.Where("user_id = ?", userId).Delete(&cart).Error; err != nil {
		return err
	}

	return nil
}

// FindByUserId implements CartRepository
func (repository *CartRepositoryImpl) FindByUserId(userId int, offset int, limit int) []entity.Cart {
	var carts []entity.Cart

	repository.db.Scopes(utils.Paginate(offset, limit)).Preload("User").Preload("Product.ProductCategory").Where("user_id = ?", userId).Find(&carts)

	return carts
}

// Create implements CartRepository
func (repository *CartRepositoryImpl) Create(entity entity.Cart) (*entity.Cart, error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// Delete implements CartRepository
func (repository *CartRepositoryImpl) Delete(entity entity.Cart) error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return err
	}

	return nil
}

// FindById implements CartRepository
func (repository *CartRepositoryImpl) FindById(id int) (*entity.Cart, error) {
	var cart entity.Cart

	if err := repository.db.Preload("User").Preload("Product").Find(&cart, id).Error; err != nil {
		return nil, err
	}

	return &cart, nil
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &CartRepositoryImpl{db}
}
