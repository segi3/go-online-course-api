package profile

import (
	dto "online-course/internal/profile/dto"
	useUseCase "online-course/internal/user/usecase"
)

type ProfileUseCase interface {
	GetProfile(id int) (*dto.ProfileResponseBody, error)
}

type ProfileUseCaseImpl struct {
	userUseCase useUseCase.UserUseCase
}

// GetProfile implements ProfileUseCase
func (usecase *ProfileUseCaseImpl) GetProfile(id int) (*dto.ProfileResponseBody, error) {
	user, err := usecase.userUseCase.FindById(id)

	if err != nil {
		return nil, err
	}

	userResponse := dto.CreateProfileResponse(*user)

	return &userResponse, nil
}

func NewProfileUseCase(userUseCase useUseCase.UserUseCase) ProfileUseCase {
	return &ProfileUseCaseImpl{userUseCase}
}
