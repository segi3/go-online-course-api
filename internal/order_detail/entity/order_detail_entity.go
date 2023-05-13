package order_detail

import (
	"database/sql"

	productEntity "online-course/internal/product/entity"
	userEntity "online-course/internal/user/entity"

	"gorm.io/gorm"
)

type OrderDetail struct {
	ID          int64                  `json:"id"`
	Price       int64                  `json:"price"`
	OrderID     int64                  `json:"order_id"`
	Product     *productEntity.Product `json:"product" gorm:"foreignKey:ProductID;references:ID"`
	ProductID   int64                  `json:"product_id"`
	CreatedByID int64                  `json:"created_by" gorm:"column:created_by"`
	CreatedBy   *userEntity.User       `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID *int64                 `json:"updated_by"  gorm:"column:updated_by"`
	UpdatedBy   *userEntity.User       `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt   sql.NullTime           `json:"created_at"`
	UpdatedAt   sql.NullTime           `json:"updated_at"`
	DeletedAt   gorm.DeletedAt         `json:"deleted_at"`
}
