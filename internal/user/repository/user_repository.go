package user

import (
	entity "online-course/internal/user/entity"
	"online-course/pkg/utils"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(offest int, limit int) []entity.User
	FindById(id int) (*entity.User, error)
	Count() int
	FindByEmail(email string) (*entity.User, error)
	Create(entity entity.User) (*entity.User, error)
	Update(entity entity.User) (*entity.User, error)
	Delete(entity entity.User) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

// Count implements UserRepository
func (repository *UserRepositoryImpl) Count() int {
	var user entity.User

	var totalUser int64

	repository.db.Model(&user).Count(&totalUser)

	return int(totalUser)
}

// FindByEmail implements UserRepository
func (repository *UserRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	if err := repository.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// FindAll implements UserRepository
func (repository *UserRepositoryImpl) FindAll(offest int, limit int) []entity.User {
	var users []entity.User

	repository.db.Scopes(utils.Paginate(offest, limit)).Find(&users)

	return users
}

// FindById implements UserRepository
func (repository *UserRepositoryImpl) FindById(id int) (*entity.User, error) {
	var user entity.User

	if err := repository.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Create implements UserRepository
func (repository *UserRepositoryImpl) Create(entity entity.User) (*entity.User, error) {

	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// Update implements UserRepository
func (repository *UserRepositoryImpl) Update(entity entity.User) (*entity.User, error) {

	if err := repository.db.Save(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// Delete implements UserRepository
func (repository *UserRepositoryImpl) Delete(entity entity.User) error {
	if err := repository.db.Delete(&entity).Error; err != nil {
		return err
	}

	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}
