package user

import (
	"errors"
	"fmt"

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
	Update(id int, userDto dto.UserRequestBody) (*entity.User, error)
	Count() int
	Delete(id int) error
}

type UserUseCaseImpl struct {
	repository repository.UserRepository
}

// Count implements UserUseCase
func (usecase *UserUseCaseImpl) Count() int {
	return usecase.repository.Count()
}

// FindByEmail implements UserUseCase
func (usecase *UserUseCaseImpl) FindByEmail(email string) (*entity.User, error) {
	return usecase.repository.FindByEmail(email)
}

// Create implements UserUseCase
func (usecase *UserUseCaseImpl) Create(userDto dto.UserRequestBody) (*entity.User, error) {
	// Find by email
	checkUser, err := usecase.repository.FindByEmail(*userDto.Email)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if checkUser != nil {
		return nil, errors.New("email sudah pernah terdaftar")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*userDto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := entity.User{
		Name:         *userDto.Name,
		Email:        *userDto.Email,
		Password:     string(hashedPassword),
		CodeVerified: utils.RandString(32),
	}

	if userDto.CreatedBy != nil {
		user.CreatedByID = userDto.CreatedBy
	}

	fmt.Println(user)
	// Create data
	dataUser, err := usecase.repository.Create(user)

	if err != nil {
		return nil, err
	}

	return dataUser, nil
}

// Delete implements UserUseCase
func (usecase *UserUseCaseImpl) Delete(id int) error {
	user, err := usecase.repository.FindById(id)

	if err != nil {
		return err
	}

	err = usecase.repository.Delete(*user)

	if err != nil {
		return err
	}

	return nil
}

// FindAll implements UserUseCase
func (usecase *UserUseCaseImpl) FindAll(offset int, limit int) []entity.User {
	return usecase.repository.FindAll(offset, limit)
}

// FindById implements UserUseCase
func (usecase *UserUseCaseImpl) FindById(id int) (*entity.User, error) {
	return usecase.repository.FindById(id)
}

// Update implements UserUseCase
func (usecase *UserUseCaseImpl) Update(id int, dto dto.UserRequestBody) (*entity.User, error) {
	user, err := usecase.repository.FindById(id)

	if err != nil {
		return nil, err
	}

	if dto.Name != nil {
		user.Name = *dto.Name
	}

	if dto.Email != nil {
		if user.Email != *dto.Email {
			user.Email = *dto.Email
		}
	}

	if dto.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*dto.Password), bcrypt.DefaultCost)

		if err != nil {
			return nil, err
		}

		user.Password = string(hashedPassword)
	}

	if dto.UpdatedBy != nil {
		user.UpdatedByID = dto.UpdatedBy
	}

	updateUser, err := usecase.repository.Update(*user)

	if err != nil {
		return nil, err
	}

	return updateUser, nil
}

func NewUserUseCase(repository repository.UserRepository) UserUseCase {
	return &UserUseCaseImpl{repository}
}
