package deposit

import (
	"context"
	"go-gin-codef-api/db"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

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

	if strings.ReplaceAll(year, " ", "") == "" || strings.ReplaceAll(month, " ", "") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid_request",
		})
		return
	}

	date := year + month
	lastDay := getLastDay(year, month)

	filter := bson.M{
		"resdepositdate": bson.M{"$gte": date + "01", "$lte": date + lastDay},
	}

	cursor, err := client.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	resultMap := map[string]interface{}{}
	calendarData, totalAmount := getDepositSum("date", cursor)

	resultMap["calendar_data"] = calendarData // 일자별 입금 금액
	resultMap["total_amount"] = totalAmount   // 총 입금 금액

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    resultMap,
	})

	cursor.Close(context.Background())
	client.Database().Client().Disconnect(context.TODO())
}

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
	cardDataMap, totalAmount := getDepositSum("card", cursor)

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
