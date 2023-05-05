package user

import (
	entity "online-course/internal/user/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(offset int, limit int) []entity.User
	FindById(id int) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Create(entity entity.User) (*entity.User, error)
	Update(entity entity.User) (*entity.User, error)
	Delete(entity entity.User) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

// FindByEmail implements UserRepository
func (ur *UserRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Findall implements UserRepository
func (ur *UserRepositoryImpl) FindAll(offset int, limit int) []entity.User {
	var users []entity.User

	ur.db.Find(&users)

	return users
}

// FindById implements UserRepository
func (ur *UserRepositoryImpl) FindById(id int) (*entity.User, error) {
	var user entity.User

	if err := ur.db.Where("id", id).First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil

}

// Create implements UserRepository
func (ur *UserRepositoryImpl) Create(entity entity.User) (*entity.User, error) {
	if err := ur.db.Create(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// Update implements UserRepository
func (ur *UserRepositoryImpl) Update(entity entity.User) (*entity.User, error) {
	if err := ur.db.Save(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// Save implements UserRepository
func (ur *UserRepositoryImpl) Delete(entity entity.User) error {
	if err := ur.db.Save(&entity).Error; err != nil {
		return err
	}

	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}
