package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/irisAlex/ai/internal/middleware"
	"github.com/irisAlex/ai/pkg/global"
	"github.com/irisAlex/ai/pkg/plugin/announcement/router"
)

func Router(engine *gin.Engine) {
	public := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("")
	private := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("")
	private.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	router.Router.Info.Init(public, private)
}
