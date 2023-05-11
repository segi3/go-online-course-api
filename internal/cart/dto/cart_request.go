package cart

type CartRequestBody struct {
	ProductID int64 `json:"product_id" binding:"required,number"`
	UserID    int64 `json:"user_id"`
}
