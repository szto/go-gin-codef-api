package utils

import (
	"context"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DEPOSIT_SUM_DATE = "date"
	DEPOSIT_SUM_CARD = "card"
)

/*
입금내역 데이터
*/
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
func GetLastDay(year, month string) string {
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
입금내역 합산
*/
func GetDepositSum(depositKind string, cursor *mongo.Cursor) (map[string]map[string]int, int) {
	dataMap := map[string]map[string]int{}
	var totalAmount int

	for cursor.Next(context.Background()) {
		var elem *Datas
		var tempKind string

		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		// 파라미터에 따른 합산 로직 분리
		switch depositKind {
		case DEPOSIT_SUM_DATE:
			tempKind = elem.ResDepositDate // 일자별
		case DEPOSIT_SUM_CARD:
			tempKind = elem.ResCardCompany // 카드사별
		}

		// 총입금 금액
		tempTotalAmount, _ := strconv.Atoi(elem.ResAccountIn)
		totalAmount += tempTotalAmount

		// 입금금액
		resAccountIn, _ := strconv.Atoi(elem.ResAccountIn)

		if dataMap[tempKind] == nil {
			tempMap := map[string]int{}
			tempMap["sum_of_capture_amount"] = resAccountIn

			dataMap[tempKind] = tempMap
		} else {
			dataMap[tempKind]["sum_of_capture_amount"] += resAccountIn
		}
	}

	return dataMap, totalAmount
}
