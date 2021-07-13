# 요약
일자별, 카드사별 입금내역을 받아오는 API 서버

# ENDPOINT
### URI : /deposit
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

### URI :  /deposit
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