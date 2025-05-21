package response

import "github.com/irisAlex/ai/internal/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
