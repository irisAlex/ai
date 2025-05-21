package service

import (
	"github.com/irisAlex/ai/internal/service/example"
	"github.com/irisAlex/ai/internal/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}
