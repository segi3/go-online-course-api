package class_room

import (
	entity "online-course/internal/class_room/entity"
	"online-course/pkg/utils"

	"gorm.io/gorm"
)

type ClassRoomRepository interface {
	FindAllByUserId(offset int, limit int, userId int) []entity.ClassRoom
	FindOneByUserIdAndProductId(userId int, productId int) (*entity.ClassRoom, error)
	Create(entity entity.ClassRoom) (*entity.ClassRoom, error)
}

type ClassRoomRepositoryImpl struct {
	db *gorm.DB
}

// Create implements ClassRoomRepository
func (repository *ClassRoomRepositoryImpl) Create(entity entity.ClassRoom) (*entity.ClassRoom, error) {
	if err := repository.db.Create(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

// FindAllByUserId implements ClassRoomRepository
func (repository *ClassRoomRepositoryImpl) FindAllByUserId(offset int, limit int, userId int) []entity.ClassRoom {
	var classRooms []entity.ClassRoom

	repository.db.Scopes(utils.Paginate(offset, limit)).
		Preload("Product.ProductCategory").
		Where("user_id = ?", userId).
		Find(&classRooms)

	return classRooms
}

// FindOneByUserIdAndProductId implements ClassRoomRepository
func (repository *ClassRoomRepositoryImpl) FindOneByUserIdAndProductId(userId int, productId int) (*entity.ClassRoom, error) {
	var classRoom entity.ClassRoom

	if err := repository.db.
		Where("user_id = ?", userId).
		Where("product_id", productId).
		First(&classRoom).Error; err != nil {
		return nil, err
	}

	return &classRoom, nil
}

func NewClassRoomRepository(db *gorm.DB) ClassRoomRepository {
	return &ClassRoomRepositoryImpl{db}
}
