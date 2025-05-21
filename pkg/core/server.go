package core

import (
	"fmt"
	"time"

	"github.com/irisAlex/ai/initialize"
	"github.com/irisAlex/ai/internal/service/system"
	"github.com/irisAlex/ai/pkg/global"
)

func RunServer() {
	// if global.GVA_CONFIG.System.UseRedis {
	// 	// 初始化redis服务
	// 	initialize.Redis()
	// 	if global.GVA_CONFIG.System.UseMultipoint {
	// 		initialize.RedisList()
	// 	}
	// }

	// if global.GVA_CONFIG.System.UseMongo {
	// 	err := initialize.Mongo.Initialization()
	// 	if err != nil {
	// 		zap.L().Error(fmt.Sprintf("%+v", err))
	// 	}
	// }
	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	initServer(address, Router, 10*time.Minute, 10*time.Minute)
}
