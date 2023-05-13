package user

import (
	entity "online-course/internal/user/entity"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

// Count implements UserRepository
func (repository *UserRepositoryMock) Count() int {
	panic("unimplemented")
}

// FindByEmail implements UserRepository
func (repository *UserRepositoryMock) FindByEmail(email string) (*entity.User, error) {
	panic("unimplemented")
}

// FindAll implements UserRepository
func (repository *UserRepositoryMock) FindAll(offest int, limit int) []entity.User {
	panic("unimplemented")

}

// FindById implements UserRepository
func (repository *UserRepositoryMock) FindById(id int) (*entity.User, error) {
	arguments := repository.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil, nil
	}

	user := arguments.Get(0).(entity.User)

	return &user, nil
}

// Create implements UserRepository
func (repository *UserRepositoryMock) Create(entity entity.User) (*entity.User, error) {
	panic("unimplemented")

}

// Update implements UserRepository
func (repository *UserRepositoryMock) Update(entity entity.User) (*entity.User, error) {
	panic("unimplemented")

}

// Delete implements UserRepository
func (repository *UserRepositoryMock) Delete(entity entity.User) error {
	panic("unimplemented")

}
