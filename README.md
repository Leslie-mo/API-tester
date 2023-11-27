# API-tester
A simulator that allows you to fill in your own desired return values in the database, and obtain specific return values based on specific keys of the API you want to test.

## 1. DB
### DB Installation (mysql)
```
sudo apt install mysql-server mysql-client
```

### Start Service (To stop: stop)
```
sudo service mysql start
```
### Connect to MySQL
```
sudo mysql -u root -p
※The initial password for the root user is blank

# Set a password for the root user
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '<password>';
FLUSH PRIVILEGES;

Update connection information in 'env/'
```

### Create Database (Schema)
```
CREATE DATABASE tester_db;
```

### Select Database
```
use tester_db
```
### Create Test Table / Insert Data
```
Execute SQL in '/sql'

Notes
- If multiple corresponding KEY_ITEMS are found in the request message, the item will be obtained based on the latest update time.
- Both SLEEP_TIME and TIMEOUT_TIME units are seconds.
- If there are multiple request items (ex: "labels": ["LabelSuccess001", "LabelSuccess002"]), please register in the RESPONSE table as 'value1, value2, ...' (ex: LabelSuccess001, LabelSuccess002)
```
## 2. Backend Startup
```
go run main.go
```
## 3. Connectivity Check

CURL:
```
curl --location 'http://localhost:8888/transaction:pay' \
--header 'Content-Type: application/json' \
--header 'Judge: test' \
--data '{
    "requestId": "requestId123",
    "labels": "LabelSuccess001"
}
'
```

