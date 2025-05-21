package response

import (
	"github.com/irisAlex/ai/internal/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
