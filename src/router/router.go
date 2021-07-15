package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-codef-api/src/controller"
)

func Router() *gin.Engine {
	r := gin.Default()

	depositGet := r.Group("/deposit")
	{
		depositGet.GET("", controller.GetDepositDailyList)
		depositGet.GET("/:date", controller.GetDepositDailyDetail)
	}

	businessGet := r.Group("/business")
	{
		businessGet.GET("/status", controller.GetBusinessStatus)
	}
	return r
}
