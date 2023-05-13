package class_room

import (
	"errors"

	dto "online-course/internal/class_room/dto"
	entity "online-course/internal/class_room/entity"
	repository "online-course/internal/class_room/repository"

	"gorm.io/gorm"
)

type ClassRoomUseCase interface {
	FindAllByUserId(offset int, limit int, userId int) dto.ClassRoomListResponse
	Create(dto dto.ClassRoom) (*entity.ClassRoom, error)
}

type ClassRoomUseCaseImpl struct {
	repository repository.ClassRoomRepository
}

// Create implements ClassRoomUseCase
func (usecase *ClassRoomUseCaseImpl) Create(dto dto.ClassRoom) (*entity.ClassRoom, error) {
	// Validasi apakah user id dan product id sudah ada
	dataClassRoom, err := usecase.repository.FindOneByUserIdAndProductId(int(dto.UserID), int(dto.ProductID))

	// Jika data tidak ada, tidak akan masuk ke dalam if ini
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if dataClassRoom != nil {
		return nil, errors.New("anda sudah masuk ke dalam class ini")
	}

	classRoom := entity.ClassRoom{
		UserID:    dto.UserID,
		ProductID: dto.ProductID,
	}

	data, err := usecase.repository.Create(classRoom)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// FindAllByUserId implements ClassRoomUseCase
func (usecase *ClassRoomUseCaseImpl) FindAllByUserId(offset int, limit int, userId int) dto.ClassRoomListResponse {
	classRooms := usecase.repository.FindAllByUserId(offset, limit, userId)

	classRoomsResp := dto.CreateClassRoomListResponse(classRooms)

	return classRoomsResp
}

func NewClassRoomUseCase(repository repository.ClassRoomRepository) ClassRoomUseCase {
	return &ClassRoomUseCaseImpl{repository}
}
