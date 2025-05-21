package ast

type Type string

func (r Type) String() string {
	return string(r)
}

func (r Type) Group() string {
	switch r {
	case TypePackageApiEnter:
		return "ApiGroup"
	case TypePackageRouterEnter:
		return "RouterGroup"
	case TypePackageServiceEnter:
		return "ServiceGroup"
	case TypePackageApiModuleEnter:
		return "ApiGroup"
	case TypePackageRouterModuleEnter:
		return "RouterGroup"
	case TypePackageServiceModuleEnter:
		return "ServiceGroup"
	case TypePluginApiEnter:
		return "api"
	case TypePluginRouterEnter:
		return "router"
	case TypePluginServiceEnter:
		return "service"
	default:
		return ""
	}
}

const (
	TypePackageApiEnter           = "PackageApiEnter"           // api/v1/enter.go
	TypePackageRouterEnter        = "PackageRouterEnter"        // internal/router/enter.go
	TypePackageServiceEnter       = "PackageServiceEnter"       // internal/service/enter.go
	TypePackageApiModuleEnter     = "PackageApiModuleEnter"     // api/v1/{package}/enter.go
	TypePackageRouterModuleEnter  = "PackageRouterModuleEnter"  // internal/router/{package}/enter.go
	TypePackageServiceModuleEnter = "PackageServiceModuleEnter" // internal/service/{package}/enter.go
	TypePackageInitializeGorm     = "PackageInitializeGorm"     // initialize/gorm_biz.go
	TypePackageInitializeRouter   = "PackageInitializeRouter"   // initialize/router_biz.go
	TypePluginGen                 = "PluginGen"                 // pkg/plugin/{package}/gen/main.go
	TypePluginApiEnter            = "PluginApiEnter"            // pkg/plugin/{package}/enter.go
	TypePluginInitializeV1        = "PluginInitializeV1"        // initialize/plugin_biz_v1.go
	TypePluginInitializeV2        = "PluginInitializeV2"        // initialize/plugin_biz_v2.go
	TypePluginRouterEnter         = "PluginRouterEnter"         // pkg/plugin/{package}/enter.go
	TypePluginServiceEnter        = "PluginServiceEnter"        // pkg/plugin/{package}/enter.go
	TypePluginInitializeApi       = "PluginInitializeApi"       // pkg/plugin/{package}/initialize/api.go
	TypePluginInitializeGorm      = "PluginInitializeGorm"      // pkg/plugin/{package}/initialize/gorm.go
	TypePluginInitializeMenu      = "PluginInitializeMenu"      // pkg/plugin/{package}/initialize/menu.go
	TypePluginInitializeRouter    = "PluginInitializeRouter"    // pkg/plugin/{package}/initialize/router.go
)
