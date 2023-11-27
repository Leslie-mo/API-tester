DELETE FROM API_INFO;
DELETE FROM KEY_INFO;
DELETE FROM RESPONSE;
DELETE FROM CALLBACK;


# API_INFO
INSERT INTO API_INFO(API_INFO_KEY,TARGET_API)
VALUES ('API001','transaction:pay');

INSERT INTO API_INFO(API_INFO_KEY,TARGET_API)
VALUES ('API002','end_user_token');

INSERT INTO API_INFO(API_INFO_KEY,TARGET_API)
VALUES ('API003','transaction:subscribe');

INSERT INTO API_INFO(API_INFO_KEY,TARGET_API)
VALUES ('API004','transaction:pay');

INSERT INTO API_INFO(API_INFO_KEY,TARGET_API)
VALUES ('API005','transaction:pay');

INSERT INTO API_INFO(API_INFO_KEY,TARGET_API)
VALUES ('API006','listTransaction');

INSERT INTO API_INFO(API_INFO_KEY,TARGET_API)
VALUES ('API007','getTransaction');

# KEY_INFO
INSERT INTO KEY_INFO(KEY_INFO_KEY,API_INFO_KEY,KEY_ITEM,ITEM_SECTION)
VALUES ('KEY001','API001','labels','BODY');

INSERT INTO KEY_INFO(KEY_INFO_KEY,API_INFO_KEY,KEY_ITEM,ITEM_SECTION)
VALUES ('KEY002','API002','','');

INSERT INTO KEY_INFO(KEY_INFO_KEY,API_INFO_KEY,KEY_ITEM,ITEM_SECTION)
VALUES ('KEY003','API003','transactionid','PATHPARAMATER');

INSERT INTO KEY_INFO(KEY_INFO_KEY,API_INFO_KEY,KEY_ITEM,ITEM_SECTION)
VALUES ('KEY004','API004','labels','BODY');

INSERT INTO KEY_INFO(KEY_INFO_KEY,API_INFO_KEY,KEY_ITEM,ITEM_SECTION)
VALUES ('KEY005','API005','amount.value','BODY');

INSERT INTO KEY_INFO(KEY_INFO_KEY,API_INFO_KEY,KEY_ITEM,ITEM_SECTION)
VALUES ('KEY006','API006','pageSize','QUERYPARAMETER');

INSERT INTO KEY_INFO(KEY_INFO_KEY,API_INFO_KEY,KEY_ITEM,ITEM_SECTION)
VALUES ('KEY007','API007','transactionid','PATHPARAMATER');

# RESPONSE
INSERT INTO RESPONSE(RESPONSE_KEY,KEY_INFO_KEY,KEY_VALUE,HTTPSTATUS,RE_TELEGRAM,SLEEP_TIME)
VALUES ('RESPONSE001','KEY001','LabelSuccess001',200,'{
    "success": true,
    "requestId": "{replace}",
    "resultCode": 200,
    "resultDescription": "Success",
    "resultProperty": {
        "property1": 121241
    },
    "transactionId": "rand(26, alphaUpperNum)",
    "status": "Completed",
    "receivedTime": "nowTime()",
    "orderId": "{replace}",
    "resultRegisterAccountProperty": "accountProperty"
}',0);

INSERT INTO RESPONSE(RESPONSE_KEY,KEY_INFO_KEY,KEY_VALUE,HTTPSTATUS,RE_TELEGRAM,SLEEP_TIME)
VALUES ('RESPONSE002','KEY002','',200,'{
    "expiresAt":  "nowTime()",
    "token": "rand(30, alphaNumSym)"
}',1);

INSERT INTO RESPONSE(RESPONSE_KEY,KEY_INFO_KEY,KEY_VALUE,HTTPSTATUS,RE_TELEGRAM,SLEEP_TIME)
VALUES ('RESPONSE003','KEY003','RMZ3PV3WOEILD4ZPZOZBYAVID',200,'{
"subscribeId": "01FC5K6EZ7Q5VTP0BHEFES3SP3"
}',0);

INSERT INTO RESPONSE(RESPONSE_KEY,KEY_INFO_KEY,KEY_VALUE,HTTPSTATUS,RE_TELEGRAM,SLEEP_TIME)
VALUES ('RESPONSE004','KEY004','LabelFail001',401,'{
  "code": 401,
  "message": "unauthorized"
}',2);

INSERT INTO RESPONSE(RESPONSE_KEY,KEY_INFO_KEY,KEY_VALUE,HTTPSTATUS,RE_TELEGRAM,SLEEP_TIME)
VALUES ('RESPONSE005','KEY005','1200',200,'{
  "code": 200,
  "message": "Json test"
}',0);

INSERT INTO RESPONSE(RESPONSE_KEY,KEY_INFO_KEY,KEY_VALUE,HTTPSTATUS,RE_TELEGRAM,SLEEP_TIME)
VALUES ('RESPONSE006','KEY006','100',200,'[
  {
    "action": "CAPTURE",
    "amount": {
      "currencyCode": "JPY",
      "value": 1200
    },
    "baseTransactionId": "01DQ4H6BA0ZPX4V3DOR7TJ0J76",
    "paymentGroupId": "01DQ4H6WA0ZPX4V3GRY7TJ0J70",
    "paymentMethodId": "PayPay",
    "relatedTransactionId": "01DQ4H6BA0ZPX4V3DOR7TJ0J76",
    "requestId": "sampleId_02",
    "requestProperty": {},
    "resultCode": 100,
    "resultDescription": "正常に処理が終了しました",
    "resultProperty": {},
    "status": "SUCCESS",
    "transactionId": "01DQ4H0BA0Z4X4V5DOR7TJ0l4",
    "labels": [
      "ラベル"
    ],
    "orderId": "order_01",
    "receivedTime": "2021-10-12T11:11:57+09:00",
    "processedTime": "2021-10-13T11:11:57+09:00"
  },
    {
    "action": "PAY",
    "amount": {
      "currencyCode": "JPY",
      "value": 1200
    },
    "baseTransactionId": "01DQ4H6BA0ZPX4V3DOR7TJ0J76",
    "paymentGroupId": "01DQ4H6WA0ZPX4V3GRY7TJ0J70",
    "paymentMethodId": "PayPay",
    "relatedTransactionId": "01DQ4H6BA0ZPX4V3DOR7TJ0J76",
    "requestId": "sampleId_02",
    "requestProperty": {},
    "resultCode": 100,
    "resultDescription": "正常に処理が終了しました",
    "resultProperty": {},
    "status": "SUCCESS",
    "transactionId": "01DQ4H0BA0Z4X4V5DOR7TJ0l4",
    "labels": [
      "ラベル"
    ],
    "orderId": "order_01",
    "receivedTime": "2021-10-12T11:11:57+09:00",
    "processedTime": "2021-10-13T11:11:57+09:00"
  }
]',0);

INSERT INTO RESPONSE(RESPONSE_KEY,KEY_INFO_KEY,KEY_VALUE,HTTPSTATUS,RE_TELEGRAM,SLEEP_TIME)
VALUES ('RESPONSE007','KEY007','01DQ4H6BA0ZPX4V3DOR7TJ0J76',200,'{
    "action": "PAY",
    "amount": {
      "currencyCode": "JPY",
      "value": 1200
    },
    "baseTransactionId": "01DQ4H6BA0ZPX4V3DOR7TJ0J76",
    "paymentGroupId": "01DQ4H6WA0ZPX4V3GRY7TJ0J70",
    "paymentMethodId": "PayPay",
    "relatedTransactionId": "01DQ4H6BA0ZPX4V3DOR7TJ0J76",
    "requestId": "sampleId_02",
    "requestProperty": {},
    "resultCode": 100,
    "resultDescription": "正常に処理が終了しました",
    "resultProperty": {},
    "status": "SUCCESS",
    "transactionId": "01DQ4H0BA0Z4X4V5DOR7TJ0l4",
    "labels": [
      "ラベル"
    ],
    "orderId": "order_01",
    "receivedTime": "2021-10-12T11:11:57+09:00",
    "processedTime": "2021-10-13T11:11:57+09:00"
  }',0);


