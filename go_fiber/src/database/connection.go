package database

import (
    "context"
    "os"
    "fmt"
    "time"
    "errors"
    "log"
 
    _ "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
)

var instance *mongo.Client

func Connect(envVar string) (*mongo.Client, context.Context, context.CancelFunc, *error) {
	// get environment var
	connectStr := os.Getenv(envVar)
	if len(connectStr) == 0 {
		err := errors.New(fmt.Sprintf("No %s defined in environment", envVar))
		return nil, nil, nil, &err
	}
	// create client
	log.Println(fmt.Sprintf("Connecting to db via: %s", connectStr))
	client, err := mongo.NewClient(options.Client().ApplyURI(connectStr))
	if err != nil {
		log.Fatal(err)
	}
	
	// connection timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// test connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to db.")

	instance = client
	return client, ctx, cancel, &err
}

/*

//GetDB ...
func GetDB() *gorp.DbMap {
	return db
}

//RedisClient ...
var RedisClient *_redis.Client

//InitRedis ...
func InitRedis(selectDB ...int) {

	var redisHost = os.Getenv("REDIS_HOST")
	var redisPassword = os.Getenv("REDIS_PASSWORD")

	RedisClient = _redis.NewClient(&_redis.Options{
		Addr:     redisHost,
		Password: redisPassword,
		DB:       selectDB[0],
		// DialTimeout:        10 * time.Second,
		// ReadTimeout:        30 * time.Second,
		// WriteTimeout:       30 * time.Second,
		// PoolSize:           10,
		// PoolTimeout:        30 * time.Second,
		// IdleTimeout:        500 * time.Millisecond,
		// IdleCheckFrequency: 500 * time.Millisecond,
		// TLSConfig: &tls.Config{
		// 	InsecureSkipVerify: true,
		// },
	})

}

//GetRedis ...
func GetRedis() *_redis.Client {
	return RedisClient
}

*/