package cart

import (
	"errors"

	dto "online-course/internal/cart/dto"
	entity "online-course/internal/cart/entity"
	repository "online-course/internal/cart/repository"
)

type CartUseCase interface {
	FindByUserId(userId int, offset int, limit int) []entity.Cart
	FindById(id int) (*entity.Cart, error)
	Create(dto dto.CartRequestBody) (*entity.Cart, error)
	Delete(id int, userId int) error
	DeleteByUserId(userId int) error
}

type CartUseCaseImpl struct {
	repository repository.CartRepository
}

// DeleteByUserId implements CartUseCase
func (usecase *CartUseCaseImpl) DeleteByUserId(userId int) error {
	err := usecase.repository.DeleteByUserId(userId)

	if err != nil {
		return err
	}

	return nil
}

// Create implements CartUseCase
func (usecase *CartUseCaseImpl) Create(dto dto.CartRequestBody) (*entity.Cart, error) {
	cart := entity.Cart{
		UserID:    dto.UserID,
		ProductID: dto.ProductID,
	}

	// Perlu validasi untuk memeriksa apakah user pernah menginput data dengan product id yang sama

	// Input data
	data, err := usecase.repository.Create(cart)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// Delete implements CartUseCase
func (usecase *CartUseCaseImpl) Delete(id int, userId int) error {
	// Search berdasarkan id
	cart, err := usecase.repository.FindById(id)

	if err != nil {
		return err
	}

	if cart.User.ID != int64(userId) {
		return errors.New("cart ini bukan milik anda")
	}

	err = usecase.repository.Delete(*cart)

	if err != nil {
		return err
	}

	return nil
}

// FindById implements CartUseCase
func (usecase *CartUseCaseImpl) FindById(id int) (*entity.Cart, error) {
	return usecase.repository.FindById(id)
}

// FindByUserId implements CartUseCase
func (usecase *CartUseCaseImpl) FindByUserId(userId int, offset int, limit int) []entity.Cart {
	return usecase.repository.FindByUserId(userId, offset, limit)
}

func NewCartUseCase(repository repository.CartRepository) CartUseCase {
	return &CartUseCaseImpl{repository}
}
