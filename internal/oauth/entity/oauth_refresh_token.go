package oauth

import "database/sql"

type OauthRefreshToken struct {
	ID                 int64             `json:"id"`
	OauthAccessToken   *OauthAccessToken `gorm:"foreignKey:OauthAccessTokenID;references:ID"`
	OauthAccessTokenID *int64            `json:"oauth_access_token_id"`
	UserID             int64             `json:"user_id"`
	Token              string            `json:"token"`
	ExpiredAt          sql.NullTime      `json:"expired_at"`
	CreatedAt          sql.NullTime      `json:"created_at"`
	UpdatedAt          sql.NullTime      `json:"updated_at"`
	DeletedAt          sql.NullTime      `json:"deleted_at"`
}
