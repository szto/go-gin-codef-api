package main

import (
	"context"
	"db"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Datas struct {
	CommEndDate          string
	CommMemberStoreGroup string
	CommStartDate        string
	ResAccountIn         string
	ResBankName          string
	ResCardCompany       string
	ResDepositDate       string
	ResMemberStoreNo     string
	ResOtherDeposit      string
	ResPaymentAccount    string
	ResSalesAmount       string
	ResSalesCount        string
	ResSuspenseAmount    string
}

/*
마지막날 계산
*/
func getLastDay(year, month string) string {
	lastDayMap := map[int]string{
		1: "31", 2: "28",
		3: "31", 4: "30",
		5: "31", 6: "30",
		7: "31", 8: "31",
		9: "30", 10: "31",
		11: "30", 12: "31",
	}

	tempYear, _ := strconv.Atoi(year)
	tempMonth, _ := strconv.Atoi(month)

	// 2월의 경우 윤년 계산
	if tempMonth == 2 && ((tempYear%4 == 0 && tempYear%100 != 0) || tempYear%400 == 0) {
		lastDayMap[tempMonth] = "29"
	}

	return lastDayMap[tempMonth]
}

/*
 일별 합계
*/
func depositDailyList(c *gin.Context) {
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

	var totalAmount int
	calendarData := map[string]map[string]int{}

	for cursor.Next(context.Background()) {
		var elem *Datas
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		// 총입금 금액
		tempTotalAmount, _ := strconv.Atoi(elem.ResAccountIn)
		totalAmount += tempTotalAmount

		// 일자별 입금 금액
		resAccountIn, _ := strconv.Atoi(elem.ResAccountIn)
		if calendarData[elem.ResDepositDate] == nil {
			tempMap := map[string]int{}
			tempMap["sum_of_capture_amount"] = resAccountIn
			calendarData[elem.ResDepositDate] = tempMap
		} else {
			calendarData[elem.ResDepositDate]["sum_of_capture_amount"] += resAccountIn
		}
	}

	resultMap := map[string]interface{}{}
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
func depositDailyDetail(c *gin.Context) {
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

	cardDataMap := map[string]map[string]int{}
	var totalAmount int

	for cursor.Next(context.Background()) {
		var elem *Datas
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		// 총입금 금액
		tempTotalAmount, _ := strconv.Atoi(elem.ResAccountIn)
		totalAmount += tempTotalAmount

		// 카드사별 입금 금액 및 매출 금액
		resAccountIn, _ := strconv.Atoi(elem.ResAccountIn) // 입금금액

		if cardDataMap[elem.ResCardCompany] == nil {
			tempMap := map[string]int{}
			tempMap["sum_of_capture_amount"] = resAccountIn

			cardDataMap[elem.ResCardCompany] = tempMap
		} else {
			cardDataMap[elem.ResCardCompany]["sum_of_capture_amount"] += resAccountIn
		}
	}

	resultMap := map[string]interface{}{}
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

func main() {
	r := gin.Default()

	depositGet := r.Group("/deposit")
	{
		depositGet.GET("", depositDailyList)
		depositGet.GET("/:date", depositDailyDetail)
	}

	r.Run()
}
