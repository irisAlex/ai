package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/irisAlex/ai/pkg/plugin/announcement"
	"github.com/irisAlex/ai/pkg/utils/plugin/v2"
)

func PluginInitV2(group *gin.Engine, plugins ...plugin.Plugin) {
	for i := 0; i < len(plugins); i++ {
		plugins[i].Register(group)
	}
}
func bizPluginV2(engine *gin.Engine) {
	PluginInitV2(engine, announcement.Plugin)
}
