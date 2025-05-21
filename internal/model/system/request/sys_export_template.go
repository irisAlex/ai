package request

import (
	"time"

	"github.com/irisAlex/ai/internal/model/common/request"
	"github.com/irisAlex/ai/internal/model/system"
)

type SysExportTemplateSearch struct {
	system.SysExportTemplate
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
