package product

type ProductRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required,number"`
}

type UpdateProductRequest struct {
	Name        string
	Description string
	Price       int
}
