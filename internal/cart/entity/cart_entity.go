package cart

import (
	"database/sql"

	productEntity "online-course/internal/product/entity"
	userEntity "online-course/internal/user/entity"
)

type Cart struct {
	ID        int64                  `json:"id"`
	User      *userEntity.User       `json:"user" gorm:"foreignKey:UserID;references:ID"`
	UserID    int64                  `json:"user_id"`
	Product   *productEntity.Product `json:"product" gorm:"foreignKey:ProductID;references:ID"`
	ProductID int64                  `json:"product_id"`
	CreatedAt sql.NullTime           `json:"created_at"`
	UpdatedAt sql.NullTime           `json:"updated_at"`
}
