package product

import (
	entity "online-course/internal/product/entity"
	"online-course/pkg/utils"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(offset int, limit int) []entity.Product
	FindById(id int) (*entity.Product, error)
	Count() int
	Create(entity entity.Product) (*entity.Product, error)
	Update(entity entity.Product) (*entity.Product, error)
	Delete(entity entity.Product) error
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

// Count implements ProductRepository
func (repository *ProductRepositoryImpl) Count() int {
	var product entity.Product

	var totalProduct int64

	repository.db.Model(&product).Count(&totalProduct)

	return int(totalProduct)
}

// Create implements ProductRepository
func (repository *ProductRepositoryImpl) Create(entity entity.Product) (*entity.Product, error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// Delete implements ProductRepository
func (repository *ProductRepositoryImpl) Delete(entity entity.Product) error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return err
	}

	return nil
}

// FindAll implements ProductRepository
func (repository *ProductRepositoryImpl) FindAll(offset int, limit int) []entity.Product {
	var products []entity.Product

	repository.db.Scopes(utils.Paginate(offset, limit)).Preload("ProductCategory").Find(&products)

	return products
}

// FindById implements ProductRepository
func (repository *ProductRepositoryImpl) FindById(id int) (*entity.Product, error) {
	var product entity.Product

	if err := repository.db.Preload("ProductCategory").First(&product, id).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

// Update implements ProductRepository
func (repository *ProductRepositoryImpl) Update(entity entity.Product) (*entity.Product, error) {
	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db}
}
