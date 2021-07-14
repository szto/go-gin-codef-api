package main

import (
	"deposit"
	"storestatus"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	depositGet := r.Group("/deposit")
	{
		depositGet.GET("", deposit.GetDepositDailyList)
		depositGet.GET("/:date", deposit.GetDepositDailyDetail)
	}

	r.GET("/storestatus", storestatus.GetCloseStoreInfo)

	r.Run()
}
