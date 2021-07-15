# 요약
 - 일자별, 카드사별 입금내역 수신 API
 - 사업자 휴폐업 상태 조회 API

# ENDPOINT
### URI : [GET] /deposit
 - description
   - 일자별 입금내역 수신
 - request
   - query string : year, month
   - 예시 : /deposit?year=2021&month=05
 - response
```
{
    "data": {
        "calendar_data": {
            "20210503": {
                "sum_of_capture_amount": 508162
            },
            "20210504": {
                "sum_of_capture_amount": 479189
            },
            "20210506": {
                "sum_of_capture_amount": 203961
            },
            "20210507": {
                "sum_of_capture_amount": 536375
            },
            "20210510": {
                "sum_of_capture_amount": 241552
            },
            "20210511": {
                "sum_of_capture_amount": 375987
            },
            "20210512": {
                "sum_of_capture_amount": 331806
            },
            "20210513": {
                "sum_of_capture_amount": 314807
            }
        },
        "total_amount": 6994606
    },
    "message": "success"
}
```

### URI : [GET] /deposit
 - description
   - 카드사별 입금내역 수신
 - request
   - path parameter : date
   - 예시 : /deposit/20210513
 - response
```
{
    "data": {
        "card_data": {
            "롯데카드": {
                "sum_of_capture_amount": 26784
            },
            "비씨카드": {
                "sum_of_capture_amount": 100396
            },
            "삼성카드": {
                "sum_of_capture_amount": 76781
            }
        },
        "search_date": "20210506",
        "total_amount": 203961
    },
    "message": "success"
}
```

###  URI : [GET] /business/status
 - description
   - 사업자번호로 휴폐업 상태 조회
 - request
   - query string : biz_number
   - 예시 : /business/status?biz_number=112-23-34455
 - response
```
{
    "data": {
        "code": "CF-00000",
        "extraMessage": "",
        "message": "성공",
        "resBusinessStatus": "부가가치세 일반과세자 입니다.",
        "resClosingDate": "",
        "resCompanyIdentityNo": "1122334455",
        "resTaxationTypeCode": "1",
        "resTransferTaxTypeDate": ""
    },
    "message": "success"
}
```