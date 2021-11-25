package routers

import (
	"trojan/distribute/routers/api"
	"trojan/distribute/routers/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Throttle(1000, 20))
	r.Use(middleware.Cors())

	transaction := r.Group("/api/bridge")
	{
		transaction.POST("/crossBsc", api.SendCrossTransaction)
		transaction.GET("/getBscBalance", api.GetBscBalance)
		transaction.GET("/getBscFee", api.GetBscFee)
		transaction.GET("/trxHash", api.GetTrxInfo)
	}

	system := r.Group("/api/system")
	{
		system.GET("/node", api.GetNodeInfo)
		system.GET("/blockNumber", api.GetBlockInfo)
	}

	return r
}
