package oauth

import "database/sql"

type OauthClient struct {
	ID           int64        `json:"id"`
	ClientID     string       `json:"client_id"`
	ClientSecret string       `json:"client_secret"`
	Name         string       `json:"name"`
	Redirect     string       `json:"redirect"`
	Description  string       `json:"description"`
	Scope        string       `json:"scope"`
	CreatedBy    int64        `json:"created_by"`
	UpdatedBy    int64        `json:"updated_by"`
	CreatedAt    sql.NullTime `json:"created_at"`
	UpdatedAt    sql.NullTime `json:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at"`
}
