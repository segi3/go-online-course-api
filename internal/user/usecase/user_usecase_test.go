package user

import (
	"testing"

	entity "online-course/internal/user/entity"
	repository "online-course/internal/user/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = &repository.UserRepositoryMock{Mock: mock.Mock{}}
var userUseCase = UserUseCaseImpl{repository: userRepository}

func TestUserUseCase_FindByIdSuccess(t *testing.T) {
	userData := entity.User{
		ID:   1,
		Name: "faerulsalamun",
	}

	userRepository.Mock.On("FindById", 1).Return(userData)

	user, err := userUseCase.FindById(1)

	assert.NotNil(t, user)
	assert.Nil(t, err)
}

func TestUserUseCase_FindByIdNotFound(t *testing.T) {
	userRepository.Mock.On("FindById", 2).Return(nil)

	user, err := userUseCase.FindById(2)

	assert.Nil(t, err)
	assert.Nil(t, user)
}
