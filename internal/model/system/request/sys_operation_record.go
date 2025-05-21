package request

import (
	"github.com/irisAlex/ai/internal/model/common/request"
	"github.com/irisAlex/ai/internal/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
