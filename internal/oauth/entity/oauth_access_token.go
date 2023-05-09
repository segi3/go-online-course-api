package oauth

import "database/sql"

type OauthAccessToken struct {
	ID            int64        `json:"id"`
	OauthClient   *OauthClient `gorm:"foreignKey:OauthClientID;references:ID"`
	OauthClientID *int64       `json:"oauth_client_id"`
	UserID        int64        `json:"user_id"`
	Token         string       `json:"token"`
	Scope         string       `json:"scope"`
	ExpiredAt     sql.NullTime `json:"expired_at"`
	CreatedAt     sql.NullTime `json:"created_at"`
	UpdatedAt     sql.NullTime `json:"updated_at"`
	DeletedAt     sql.NullTime `json:"deleted_at"`
}
