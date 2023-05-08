package user

import (
	"errors"
	dto "online-course/internal/user/dto"
	entity "online-course/internal/user/entity"
	repository "online-course/internal/user/repository"
	utils "online-course/pkg/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUseCase interface {
	FindAll(offset int, limit int) []entity.User
	FindById(id int) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Create(userDto dto.UserRequestBody) (*entity.User, error)
	Update(userDto dto.UserRequestBody) (*entity.User, error)
	Delete(id int) error
}

type UserUseCaseImpl struct {
	repository repository.UserRepository
}

// FindByEmail implements UserUseCase
func (usecase *UserUseCaseImpl) FindByEmail(email string) (*entity.User, error) {
	panic("unimplemented")
}

// Create implements UserUseCase
func (uc *UserUseCaseImpl) Create(userDto dto.UserRequestBody) (*entity.User, error) {

	// Find by email
	userExist, err := uc.repository.FindByEmail(*userDto.Email)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // make sure not found record is not catched
		return nil, err
	}

	if userExist != nil {
		return nil, errors.New("Email already taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*userDto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := entity.User{
		Name:         *userDto.Name,
		Email:        *userDto.Email,
		Password:     string(hashedPassword),
		CodeVerified: utils.RandomString(12),
	}

	// create new user data
	dataUser, err := uc.repository.Create(user)

	if err != nil {
		return nil, err
	}

	return dataUser, nil
}

// Delete implements UserUseCase
func (*UserUseCaseImpl) Delete(id int) error {
	panic("unimplemented")
}

// FindAll implements UserUseCase
func (*UserUseCaseImpl) FindAll(offset int, limit int) []entity.User {
	panic("unimplemented")
}

// FindById implements UserUseCase
func (usecase *UserUseCaseImpl) FindById(id int) (*entity.User, error) {
	return usecase.repository.FindById(id)
}

// Update implements UserUseCase
func (*UserUseCaseImpl) Update(userDto dto.UserRequestBody) (*entity.User, error) {
	panic("unimplemented")
}

func NewUserUseCase(repository repository.UserRepository) UserUseCase {
	return &UserUseCaseImpl{repository}
}
