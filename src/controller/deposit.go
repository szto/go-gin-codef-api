package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-gin-codef-api/src/db"
	"go-gin-codef-api/src/service"
	"go-gin-codef-api/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"strings"
)

/*
 카드사별 합계
*/
func GetDepositDailyDetail(c *gin.Context) {
	client, err := db.ConnectDB("testStore")

	if err != nil {
		log.Fatal(err)
	}

	date := c.Param("date")

	if strings.ReplaceAll(date, " ", "") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid_request",
		})
		return
	}

	filter := bson.M{
		"resdepositdate": date,
	}

	cursor, err := client.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	resultMap := map[string]interface{}{}
	cardDataMap, totalAmount := utils.GetDepositSum("card", cursor)

	resultMap["card_data"] = cardDataMap    // 카드사별 입금 금액
	resultMap["total_amount"] = totalAmount // 총 입금 금액
	resultMap["search_date"] = date         // 조회 일자

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    resultMap,
	})

	cursor.Close(context.Background())
	client.Database().Client().Disconnect(context.TODO())
}

/*
일별 합계
*/
func GetDepositDailyList(c *gin.Context) {
	client, err := db.ConnectDB("testStore")
	if err != nil {
		log.Fatal(err)
	}

	year := c.Query("year")
	month := c.Query("month")

	if utils.IsEmptyQueryParmas(year) || utils.IsEmptyQueryParmas(month) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid_request",
		})
		return
	}

	date := utils.GetDateByYearAndMonth(year, month)
	lastDay := utils.GetLastDay(year, month)

	filter := bson.M{
		"resdepositdate": bson.M{"$gte": date + "01", "$lte": date + lastDay},
	}

	cursor, err := service.FindDepositByFilter(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "db error.",
		})
		return
	}

	resultMap := map[string]interface{}{}
	calendarData, totalAmount := utils.GetDepositSum("date", cursor)

	resultMap["calendar_data"] = calendarData // 일자별 입금 금액
	resultMap["total_amount"] = totalAmount   // 총 입금 금액

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    resultMap,
	})

	cursor.Close(context.Background())
	client.Database().Client().Disconnect(context.TODO())
}
