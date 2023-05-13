package admin

import (
	dto "online-course/internal/admin/dto"
	entity "online-course/internal/admin/entity"
	repository "online-course/internal/admin/repository"

	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase interface {
	FindAll(offset int, limit int) []entity.Admin
	FindById(id int) (*entity.Admin, error)
	FindByEmail(email string) (*entity.Admin, error)
	Count() int
	Create(dto dto.AdminRequestBody) (*entity.Admin, error)
	Update(id int, dto dto.AdminRequestBody) (*entity.Admin, error)
	Delete(id int) error
}

type AdminUseCaseImpl struct {
	repository repository.AdminRepository
}

// Count implements AdminUseCase
func (usecase *AdminUseCaseImpl) Count() int {
	return usecase.repository.Count()
}

// Create implements AdminUseCase
func (usecase *AdminUseCaseImpl) Create(dto dto.AdminRequestBody) (*entity.Admin, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*dto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	dataAdmin := entity.Admin{
		Email:       dto.Email,
		Name:        dto.Name,
		Password:    string(hashedPassword),
		CreatedByID: dto.CreatedBy,
	}

	admin, err := usecase.repository.Create(dataAdmin)

	if err != nil {
		return nil, err
	}

	return admin, nil

}

// Delete implements AdminUseCase
func (usecase *AdminUseCaseImpl) Delete(id int) error {
	// Search dari table admin berdasarkan id
	admin, err := usecase.repository.FindById(id)

	if err != nil {
		return err
	}

	if err := usecase.repository.Delete(*admin); err != nil {
		return err
	}

	return nil
}

// FindAll implements AdminUseCase
func (usecase *AdminUseCaseImpl) FindAll(offset int, limit int) []entity.Admin {
	return usecase.repository.FindAll(offset, limit)
}

// FindByEmail implements AdminUseCase
func (usecase *AdminUseCaseImpl) FindByEmail(email string) (*entity.Admin, error) {
	return usecase.repository.FindByEmail(email)
}

// FindById implements AdminUseCase
func (usecase *AdminUseCaseImpl) FindById(id int) (*entity.Admin, error) {
	return usecase.repository.FindById(id)
}

// Update implements AdminUseCase
func (usecase *AdminUseCaseImpl) Update(id int, dto dto.AdminRequestBody) (*entity.Admin, error) {
	// Search dari table admin berdasarkan id
	admin, err := usecase.repository.FindById(id)

	admin.Name = dto.Name

	// Validasi admin email jika tidak sama maka akan di update
	if admin.Email != dto.Email {
		admin.Email = dto.Email
	}

	if dto.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*dto.Password), bcrypt.DefaultCost)

		if err != nil {
			return nil, err
		}

		admin.Password = string(hashedPassword)
	}

	if err != nil {
		return nil, err
	}

	admin.UpdatedByID = &dto.UpdatedBy

	// Kita akan melakukan update dengan memanggil repository
	updateAdmin, err := usecase.repository.Update(*admin)

	if err != nil {
		return nil, err
	}

	return updateAdmin, nil
}

func NewAdminUseCase(repository repository.AdminRepository) AdminUseCase {
	return &AdminUseCaseImpl{repository}
}
