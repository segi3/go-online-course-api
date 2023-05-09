package product_category

type ProductCategoryRequestBody struct {
	Name      string `json:"name" binding:"required"`
	CreatedBy int64  `json:"created_by"`
	UpdatedBy int64  `json:"updated_by"`
}
