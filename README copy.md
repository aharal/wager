HOW TO RUN : 
1) UnZip Folder and  put code inside `/home/[USERNAME]/go/src/github.com/` 
2) SET GOPATH 
	>> export GOPATH=$HOME/go
3) Create Build and run Binary
	>> go run main.go 
4) Log Details presnt into following file :
	Logfile: /home/[USERNAME]/go/src/github.com/aharal/wager/logs/

Where, modification needs to be done for mysql db, username and password.
	/home/[USERNAME]/go/src/github.com/aharal/wager/configs/config.toml

The database to be created.
	/home/[USERNAME]/go/src/github.com/aharal/wager/database/migrations/1_infinity.up.sql

The endpoint details:
Postman collection:
	/home/[USERNAME]/go/src/github.com/aharal/wager/postman/krazybee.postman_collection.json
API :
https://localhost:8080/albums
https://localhost:8080/photos?albumId=1
http://localhost:8080/search
         - type=album&id=1
        -  type=photo&album=1&id=1


##The Structure
Infinity means(Anant)(My Own Framework) is based on clean architecture, the project structure is like below.

.
├── Makefile
├── README.md
├── api
│   ├── models
│   │   ├── models Files
│   ├── v1
│   │   ├── handles
│   │   ├── middleware
│   │   ├── repository
│   │   ├── services
│   │   ├── responses.go
│   ├── router.go
├── configs
│   ├── config.go
│   ├── mconfig.toml
│   ├── logs.go
│   ├── Other Constant / static files
├── database
│   ├── connection
│   │   ├── connection.go
│   │   ├── connectionInterface.go
│   ├── migrations
│   │   ├── vesion_databasename.up.sql
│   │   ├── vesion_databasename.down.sql
│   ├── dbMigrate.go
├── logs
├── postman(optional)
├── vendor (GoDep)
│   ├── vendor packages


----------------------------------------------------




