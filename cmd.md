# Compose

docker-compose up -d --build api_go_fiber
docker-compose up -d db_redis db_mongo db_mongo_express

# Environment

export PATH=$PATH:$GOROOT/bin
export PATH=$PATH:/usr/local/go/bin
export MONGODB_CONNECTION_STRING=mongodb://root:admin@localhost:27017
export REDIS_CONNECTION_STRING=db_redis:6379
export MYSQL_DATABASE_ACCOUNTS="go_fiber:1234,lorien:2345,felgrom:5678"

# Running & Building

go run main.go
go build go_fiber/src/main.go
