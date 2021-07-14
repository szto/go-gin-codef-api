package storestatus

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

const CLOSE_STORE_END_POINT = "/v1/kr/public/nt/business/status"
const TYPE_DEMO = 1
const CODEF_SUCCESS_CODE = "CF-00000"

func GetCloseStoreInfo(c *gin.Context) {
	var datas Datas

	identity := c.Query("identity")

	if strings.ReplaceAll(identity, " ", "") == "" {
		datas.Data = []map[string]string{}
		c.JSON(http.StatusOK, gin.H{
			"message": "invalid_request",
			"data":    datas.Data,
		})
		return
	}

	codef := codef.GetCodef()                    // codef instance 받아오기
	identityList := strings.Split(identity, ",") // 파라미터가 다건일 경우 분리
	reqData := []map[string]string{}

	// codef 요청 데이터 생성
	for _, v := range identityList {
		tempMap := map[string]string{}
		tempMap["reqIdentity"] = v
		reqData = append(reqData, tempMap)
	}

	parameter := map[string]interface{}{
		"organization":    "0004",
		"reqIdentityList": reqData,
	}

	codefResult, err := codef.RequestProduct(CLOSE_STORE_END_POINT, TYPE_DEMO, parameter)
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

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    datas.Data,
	})
}
