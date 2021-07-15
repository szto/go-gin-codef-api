package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-gin-codef-api/src/codef"
	"go-gin-codef-api/src/http/response"
	"log"
	"net/http"
	"strings"
)

const ORGANIZATION_CODE = "0004"
const BUSINESS_STATUS_END_POINT = "/v1/kr/public/nt/business/status"
const TYPE_DEMO = 1
const CODEF_SUCCESS_CODE = "CF-00000"

func GetBusinessStatus(c *gin.Context) {
	var datas codef.CodefDatas

	bizNumber := c.Query("biz_number")
	bizNumber = strings.ReplaceAll(bizNumber, "-", "")

	if strings.ReplaceAll(bizNumber, " ", "") == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid_request",
			"data":    map[string]string{},
		})
		return
	}

	codef := codef.GetCodef() // codef instance 받아오기

	reqData := []map[string]string{
		{"reqIdentity": bizNumber},
	}

	parameter := map[string]interface{}{
		"organization":    ORGANIZATION_CODE,
		"reqIdentityList": reqData,
	}

	codefResult, err := codef.RequestProduct(BUSINESS_STATUS_END_POINT, TYPE_DEMO, parameter)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(codefResult), &datas)

	if datas.Result["code"] != CODEF_SUCCESS_CODE {
		errorMsg := response.GenerateErrorMsg(datas.Result)

		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": errorMsg,
			"data":    map[string]string{},
		})
		return
	}

	// index error 방지용
	if datas.Data == nil || len(datas.Data) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "result_empty",
			"data":    map[string]string{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    datas.Data[0],
	})
}
