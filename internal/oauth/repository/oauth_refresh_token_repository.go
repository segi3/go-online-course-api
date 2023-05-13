package oauth

import (
	entity "online-course/internal/oauth/entity"

	"gorm.io/gorm"
)

type OauthRefreshTokenRepository interface {
	Create(oauthRefreshToken entity.OauthRefreshToken) (*entity.OauthRefreshToken, error)
	FindOneByToken(token string) (*entity.OauthRefreshToken, error)
	Delete(id int) error
}

type OauthRefreshTokenRepositoryImpl struct {
	db *gorm.DB
}

// Create implements OauthRefreshTokenRepository
func (repository *OauthRefreshTokenRepositoryImpl) Create(oauthRefreshToken entity.OauthRefreshToken) (*entity.OauthRefreshToken, error) {
	if result := repository.db.Create(&oauthRefreshToken).Error; result != nil {
		return nil, result
	}

	return &oauthRefreshToken, nil
}

// Delete implements OauthRefreshTokenRepository
func (repository *OauthRefreshTokenRepositoryImpl) Delete(id int) error {
	var oauthRefreshToken entity.OauthRefreshToken

	if err := repository.db.Delete(&oauthRefreshToken, id).Error; err != nil {
		return err
	}

	return nil
}

// FindOneByToken implements OauthRefreshTokenRepository
func (repository *OauthRefreshTokenRepositoryImpl) FindOneByToken(token string) (*entity.OauthRefreshToken, error) {
	var oauthRefreshToken entity.OauthRefreshToken

	if err := repository.db.Where("token = ?", token).First(&oauthRefreshToken).Error; err != nil {
		return nil, err
	}

	return &oauthRefreshToken, nil
}

func NewOauthRefreshTokenRepository(db *gorm.DB) OauthRefreshTokenRepository {
	return &OauthRefreshTokenRepositoryImpl{db}
}
