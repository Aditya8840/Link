package databases

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Aditya8840/Link/types"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/redis/go-redis/v9"
)

type manager struct {
	client *mongo.Client
	redis_client *redis.Client
	ctx context.Context
	cancel context.CancelFunc
	maxCacheSize int64
}

var Mgr Manager

type Manager interface {
	Insert(*types.URL) error
	GetOriginalURL(string) (string, error)
	evictIfNeeded() error
}

func connect() {
	err := godotenv.Load(".env")
	if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }

	MONGO_URI := os.Getenv("MONGO_URI")

	clientOption := options.Client().ApplyURI(MONGO_URI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	
	
	client, err := mongo.Connect(ctx, clientOption)

	if err != nil {
		panic(err)
	}

	redis_client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
        Password: os.Getenv("REDIS_PASSWD"),
        DB:       0,  // use default DB
	})

	err = redis_client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	Mgr = &manager{client, redis_client, ctx, cancel, 10000}
}