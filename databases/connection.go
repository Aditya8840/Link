package databases

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Aditya8840/Link/types"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type manager struct {
	client *mongo.Client
	redis_client *redis.Client
	protected_redis *redis.Client
}

var Mgr Manager

type Manager interface {
	Insert(*types.URL) error
	GetOriginalURL(string) (string, error)
	GetAndIncCounter() (int64, error)
}

func Connect() {
	err := godotenv.Load(".env")
	if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }

	MONGO_URI := os.Getenv("MONGO_URI")

	if MONGO_URI == "" {
        log.Fatalf("MONGO_URI environment variable not set")
    }
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOption := options.Client().ApplyURI(MONGO_URI).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		panic(err)
	}

	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
	panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	redisAddr := os.Getenv("REDIS_ADDR")
    if redisAddr == "" {
        redisAddr = "localhost:6379"
    }

    redis_client := redis.NewClient(&redis.Options{
        Addr:     redisAddr,
        Password: "",
        DB:       0,
    })

    protected_redis := redis.NewClient(&redis.Options{
        Addr:     redisAddr,
        Password: "",
        DB:       1,
    })

	ctx := context.Background()
    if err := redis_client.Ping(ctx).Err(); err != nil {
		panic(err)
	}

	err = redis_client.ConfigSet(ctx, "maxmemory", "1024mb").Err()
    if err != nil {
        panic(err)
    }

    err = redis_client.ConfigSet(ctx, "maxmemory-policy", "allkeys-lru").Err()
    if err != nil {
        panic(err)
    }

    fmt.Println("AllKeys-LRU eviction policy set successfully!")

	Mgr = &manager{client, redis_client, protected_redis}
}