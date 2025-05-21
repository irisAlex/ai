package request

import (
	"github.com/irisAlex/ai/internal/model/common/request"
	"github.com/irisAlex/ai/internal/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
