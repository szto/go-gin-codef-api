package main

import (
	"deposit"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	depositGet := r.Group("/deposit")
	{
		depositGet.GET("", deposit.GetDepositDailyList)
		depositGet.GET("/:date", deposit.GetDepositDailyDetail)
	}

	r.Run()
}
