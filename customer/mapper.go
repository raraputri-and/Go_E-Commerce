package customer

func ConvertToCustomerResponse(c Customer) CustomerResponse {
	return CustomerResponse {
		Name: c.Name,
	}
}