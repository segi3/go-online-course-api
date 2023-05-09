package product_category

import (
	dto "online-course/internal/product_category/dto"
	entity "online-course/internal/product_category/entity"
	repository "online-course/internal/product_category/repository"
)

type ProductCategoryUseCase interface {
	FindAll(offset int, limit int) []entity.ProductCategory
	FindById(id int) (*entity.ProductCategory, error)
	Create(dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, error)
	Update(id int, dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, error)
	Delete(id int) error
}

type ProductCategoryUseCaseImpl struct {
	repository repository.ProductCategoryRepository
}

// Create implements ProductCategoryUseCase
func (usecase *ProductCategoryUseCaseImpl) Create(dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, error) {
	productCategoryEntity := entity.ProductCategory{
		Name:        dto.Name,
		CreatedByID: dto.CreatedBy,
	}

	productCategory, err := usecase.repository.Create(productCategoryEntity)

	if err != nil {
		return nil, err
	}

	return productCategory, nil
}

// Delete implements ProductCategoryUseCase
func (usecase *ProductCategoryUseCaseImpl) Delete(id int) error {
	productCategory, err := usecase.repository.FindById(id)

	if err != nil {
		return err
	}

	if err := usecase.repository.Delete(*productCategory); err != nil {
		return err
	}

	return nil
}

// FindAll implements ProductCategoryUseCase
func (usecase *ProductCategoryUseCaseImpl) FindAll(offset int, limit int) []entity.ProductCategory {
	return usecase.repository.FindAll(offset, limit)
}

// FindById implements ProductCategoryUseCase
func (usecase *ProductCategoryUseCaseImpl) FindById(id int) (*entity.ProductCategory, error) {
	return usecase.repository.FindById(id)
}

// Update implements ProductCategoryUseCase
func (usecase *ProductCategoryUseCaseImpl) Update(id int, dto dto.ProductCategoryRequestBody) (*entity.ProductCategory, error) {
	// Cari data product category berdasarkan id
	productCategory, err := usecase.repository.FindById(id)

	if err != nil {
		return nil, err
	}

	productCategory.Name = dto.Name
	productCategory.UpdatedByID = &dto.UpdatedBy

	// Memanggil repository untuk melakukan update
	updateProductCategory, err := usecase.repository.Update(*productCategory)

	if err != nil {
		return nil, err
	}

	return updateProductCategory, nil
}

func NewProductCategoryUseCase(repository repository.ProductCategoryRepository) ProductCategoryUseCase {
	return &ProductCategoryUseCaseImpl{repository}
}
