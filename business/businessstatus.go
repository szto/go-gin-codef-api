package business

import (
	"codef"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Datas struct {
	Data   []map[string]string
	Result map[string]string
}

const BUSINESS_STATUS_END_POINT = "/v1/kr/public/nt/business/status"
const TYPE_DEMO = 1
const CODEF_SUCCESS_CODE = "CF-00000"
const ORGANIZATION_CODE = "0004"

func GetBusinessStatus(c *gin.Context) {
	var datas Datas

	bizNumber := c.Query("biz_number")
	bizNumber = strings.ReplaceAll(bizNumber, "-", "")

	if strings.ReplaceAll(bizNumber, " ", "") == "" {
		datas.Data = []map[string]string{}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid_request",
			"data":    datas.Data,
		})
		return
	}

	codef := codef.GetCodef() // codef instance 받아오기
	reqData := []map[string]string{}

	tempMap := map[string]string{}
	tempMap["reqIdentity"] = bizNumber

	reqData = append(reqData, tempMap)

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
		errorMsg := "codef_error : " +
			datas.Result["code"] + ", " +
			datas.Result["message"]

		datas.Data = []map[string]string{}
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"message": errorMsg,
			"data":    datas.Data,
		})
		return
	}

	// index error 방지용
	if len(datas.Data) <= 0 {
		datas.Data = []map[string]string{}
		c.JSON(http.StatusOK, gin.H{
			"message": "result_empty",
			"data":    datas.Data,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    datas.Data[0],
	})
}
