package router

import (
	"github.com/irisAlex/ai/internal/router/example"
	"github.com/irisAlex/ai/internal/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}
