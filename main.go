package main

import (
	business "go-gin-codef-api/business"
	deposit "go-gin-codef-api/deposit"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	depositGet := r.Group("/deposit")
	{
		depositGet.GET("", deposit.GetDepositDailyList)
		depositGet.GET("/:date", deposit.GetDepositDailyDetail)
	}

	businessGet := r.Group("/business")
	{
		businessGet.GET("/status", business.GetBusinessStatus)
	}

	r.Run()
}
