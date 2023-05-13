package product

import (
	"database/sql"

	adminEntity "online-course/internal/admin/entity"
	entity "online-course/internal/product/entity"
	productCategoryEntity "online-course/internal/product_category/entity"

	"gorm.io/gorm"
)

type ProductResponseBody struct {
	ID              int64                                  `json:"id"`
	ProductCategory *productCategoryEntity.ProductCategory `json:"product_category"`
	Title           string                                 `json:"title"`
	Image           string                                 `json:"image"`
	Video           string                                 `json:"video"`
	Description     string                                 `json:"description"`
	Price           int64                                  `json:"price"`
	CreatedBy       *adminEntity.Admin                     `json:"created_by"`
	UpdatedBy       *adminEntity.Admin                     `json:"updated_by"`
	CreatedAt       sql.NullTime                           `json:"created_at"`
	UpdatedAt       sql.NullTime                           `json:"updated_at"`
	DeletedAt       gorm.DeletedAt                         `json:"deleted_at"`
}

func CreateProductResponse(product entity.Product) ProductResponseBody {
	return ProductResponseBody{
		ProductCategory: product.ProductCategory,
		Title:           product.Title,
		Image:           *product.Image,
		Video:           *product.Video,
		Description:     product.Description,
		Price:           product.Price,
		CreatedBy:       product.CreatedBy,
		UpdatedBy:       product.UpdatedBy,
		CreatedAt:       product.CreatedAt,
		UpdatedAt:       product.UpdatedAt,
		DeletedAt:       product.DeletedAt,
	}
}

type ProductListResponse []ProductResponseBody

func CreateProductListResponse(products []entity.Product) ProductListResponse {
	productResp := ProductListResponse{}

	for _, p := range products {

		product := CreateProductResponse(p)
		productResp = append(productResp, product)
	}

	return productResp
}
