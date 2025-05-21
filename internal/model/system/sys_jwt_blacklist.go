package system

import (
	"github.com/irisAlex/ai/pkg/global"
)

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
