# Compose

docker-compose up -d --build api_go_fiber
docker-compose up -d db_redis db_mongo db_mongo_express

# Running & Building

export MONGODB_CONNECTION_STRING=mongodb://root:admin@localhost:27017
export REDIS_CONNECTION_STRING=db_redis:6379
go run main.go
go build go_fiber/src/main.go
