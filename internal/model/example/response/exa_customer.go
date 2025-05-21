package response

import "github.com/irisAlex/ai/internal/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
