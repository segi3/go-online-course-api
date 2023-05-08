package oauth

import (
	entity "online-course/internal/oauth/entity"

	"gorm.io/gorm"
)

type OauthRefreshTokenRepository interface {
	Create(oauthRefreshToken entity.OauthRefreshToken) (*entity.OauthRefreshToken, error)
	FindOneByRefreshToken(refreshToken string) (*entity.OauthRefreshToken, error)
	Delete(id int) error
}

type OauthRefreshTokenImpl struct {
	db *gorm.DB
}

// Create implements OauthRefreshTokenRepository
func (repository *OauthRefreshTokenImpl) Create(oauthRefreshToken entity.OauthRefreshToken) (*entity.OauthRefreshToken, error) {
	if err := repository.db.Create(oauthRefreshToken).Error; err != nil {
		return nil, err
	}

	return &oauthRefreshToken, nil
}

// Delete implements OauthRefreshTokenRepository
func (repository *OauthRefreshTokenImpl) Delete(id int) error {

	var oauthRefreshToken entity.OauthRefreshToken

	if err := repository.db.Delete(&oauthRefreshToken, id).Error; err != nil {
		return err
	}
	return nil
}

// FindOneByRefreshToken implements OauthRefreshTokenRepository
func (repository *OauthRefreshTokenImpl) FindOneByRefreshToken(refreshToken string) (*entity.OauthRefreshToken, error) {
	var oauthRefreshToken entity.OauthRefreshToken

	if err := repository.db.Where("token = ?", refreshToken).First(&oauthRefreshToken).Error; err != nil {
		return nil, err
	}

	return &oauthRefreshToken, nil
}

func NewOauthRefreshToken(db *gorm.DB) OauthRefreshTokenRepository {
	return &OauthRefreshTokenImpl{db}
}
