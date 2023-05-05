package user

import "database/sql"

type User struct {
	ID              int64        `json:"id"`
	Name            string       `json:"name"`
	Email           string       `json:"email"`
	Password        string       `json:"-"` // hide password, dont show in entity
	CodeVerified    string       `json:"-"`
	EmailVerifiedAt sql.NullTime `json:"email_verified_at"`
	CreatedBy       *int64       `json:"created_by"`
	UpdatedBy       *int64       `json:"updated_by"`
	CreatedAt       sql.NullTime `json:"created_at"`
	UpdatedAt       sql.NullTime `json:"updated_at"`
	DeletedAt       sql.NullTime `json:"deleted_at"`
}
