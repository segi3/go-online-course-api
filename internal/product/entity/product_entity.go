package product

import (
	"database/sql"

	adminEntity "online-course/internal/admin/entity"
	productCategoryEntity "online-course/internal/product_category/entity"

	"gorm.io/gorm"
)

type Product struct {
	ID                int64                                  `json:"id"`
	ProductCategory   *productCategoryEntity.ProductCategory `json:"product_category" gorm:"foreignKey:ProductCategoryID;references:ID"`
	ProductCategoryID int64                                  `json:"product_category_id"`
	Title             string                                 `json:"title"`
	Image             *string                                `json:"image"`
	Video             *string                                `json:"video"`
	VideoURL          *string                                `json:"video_url,omitempty" gorm:"->"`
	Description       string                                 `json:"description"`
	Price             int64                                  `json:"price"`
	CreatedByID       int64                                  `json:"created_by" gorm:"column:created_by"`
	CreatedBy         *adminEntity.Admin                     `json:"-" gorm:"foreignKey:CreatedByID;references:ID"`
	UpdatedByID       *int64                                 `json:"updated_by"  gorm:"column:updated_by"`
	UpdatedBy         *adminEntity.Admin                     `json:"-" gorm:"foreignKey:UpdatedByID;references:ID"`
	CreatedAt         sql.NullTime                           `json:"created_at"`
	UpdatedAt         sql.NullTime                           `json:"updated_at"`
	DeletedAt         gorm.DeletedAt                         `json:"deleted_at"`
}
