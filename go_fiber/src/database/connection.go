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
    "github.com/go-redis/redis"
)

var instanceMongo *mongo.Client
var instanceRedis *redis.Client

func ConnectMongo(envVar string) (*mongo.Client, context.Context, context.CancelFunc, *error) {
	// get environment var
	connectStr := os.Getenv(envVar)
	if len(connectStr) == 0 {
		err := errors.New(fmt.Sprintf("No %s defined in environment", envVar))
		return nil, nil, nil, &err
	}
	// create client
	log.Println(fmt.Sprintf("Connecting to mongodb via: %s", connectStr))
	instanceMongo, err := mongo.NewClient(options.Client().ApplyURI(connectStr))
	if err != nil {
		log.Fatal(err)
		return nil, nil, nil, &err
	}
	
	// connection timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connect
	err = instanceMongo.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, nil, nil, &err
	}

	// test connection
	err = instanceMongo.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
		return nil, nil, nil, &err
	}

	log.Println("Successfully connected to mongodb")
	return instanceMongo, ctx, cancel, &err
}

func ConnectRedis(envVar string) (*redis.Client, context.Context, context.CancelFunc, *error) {
	// get environment var
	address := os.Getenv(envVar)
	log.Print(address);
	if len(address) == 0 {
		err := errors.New(fmt.Sprintf("No %s defined in environment", envVar))
		return nil, nil, nil, &err
	}
	
	// connection timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	// create client
	log.Println(fmt.Sprintf("Connecting to redis via: %s", address))
	instanceRedis = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
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

	// test connection
	pong, err := instanceRedis.Ping().Result()
	if err != nil {
		log.Fatal(err)
		return nil, nil, nil, &err
	}
	if pong == "PONG" {
		log.Print("Successfully connected to redis")
	}
	return instanceRedis, ctx, cancel, nil;
}