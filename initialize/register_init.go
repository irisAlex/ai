package initialize

import (
	_ "github.com/irisAlex/ai/internal/source/example"
	_ "github.com/irisAlex/ai/internal/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
