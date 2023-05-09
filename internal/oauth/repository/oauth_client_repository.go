package oauth

import (
	oauth "online-course/internal/oauth/entity"

	"gorm.io/gorm"
)

type OauthClientRepository interface {
	FindByClientIdAndClientSecret(clientId string, clientSecret string) (*oauth.OauthClient, error)
}

type OauthClientRepositoryImpl struct {
	db *gorm.DB
}

// FindByClientIdAndClientSecret implements OauthClientRepository
func (repository *OauthClientRepositoryImpl) FindByClientIdAndClientSecret(clientId string, clientSecret string) (*oauth.OauthClient, error) {
	var oauthClient oauth.OauthClient

	if err := repository.db.Where("client_id = ?", clientId).Where("client_secret = ?", clientSecret).First(&oauthClient).Error; err != nil {
		return nil, err
	}

	return &oauthClient, nil
}

func NewOauthClientRepository(db *gorm.DB) OauthClientRepository {
	return &OauthClientRepositoryImpl{db}
}
