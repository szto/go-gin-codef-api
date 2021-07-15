package business

import (
	"encoding/json"
	codef "go-gin-codef-api/codef"
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

func generateErrorMsg(errmsgMap map[string]string) string {
	msg := "codef_error : " +
		errmsgMap["code"] + ", " +
		errmsgMap["message"]

	return msg
}

func GetBusinessStatus(c *gin.Context) {
	var datas Datas

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
		errorMsg := generateErrorMsg(datas.Result)

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
