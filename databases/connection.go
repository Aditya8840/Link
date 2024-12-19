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
)

type manager struct {
	client *mongo.Client
	ctx context.Context
	cancel context.CancelFunc
}

var Mgr Manager

type Manager interface {
	Insert(interface{}, string) (interface{}, error)
	GetOriginalURL(string, string) (types.URL, error)
}

func connect() *mongo.Client{
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
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	Mgr = &manager{client, ctx, cancel}
}