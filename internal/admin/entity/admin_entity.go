package admin

import (
	"database/sql"

	"gorm.io/gorm"
)

type Admin struct {
	ID          int64          `json:"id"`
	Email       string         `json:"email"`
	Name        string         `json:"name"`
	Password    string         `json:"-"`
	CreatedByID int64          `json:"created_by" gorm:"column:created_by"`
	CreatedBy   *Admin         `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID *int64         `json:"updated_by"  gorm:"column:updated_by"`
	UpdatedBy   *Admin         `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
