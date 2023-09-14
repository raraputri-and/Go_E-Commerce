package customer

type CustomerRequest struct {
	Name string `binding:"required"`
}