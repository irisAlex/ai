package response

import "github.com/irisAlex/ai/internal/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
