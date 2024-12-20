package databases

import (
	"context"
	"fmt"
	"time"

	"github.com/Aditya8840/Link/constant"
	"github.com/Aditya8840/Link/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (mgr *manager) Insert(data *types.URL) error {
	instance := mgr.client.Database(constant.DATABASE).Collection(constant.COLLECTION_NAME)
	data.ID = primitive.NewObjectID()
	_, err := instance.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}

	err = mgr.redis_client.Set(
        context.TODO(),
        data.URLCode,
        data.LongURL,
        48*time.Hour,
    ).Err()

	return err
}

func (mgr *manager) GetOriginalURL(code string) (string, error) {
	longUrl, err := mgr.redis_client.Get(context.TODO(), code).Result()
	fmt.Printf("%s", longUrl)
	if err == nil {
		return longUrl, nil
	}
	instance := mgr.client.Database(constant.DATABASE).Collection(constant.COLLECTION_NAME)
	result := types.URL{}
	err = instance.FindOne(context.TODO(), map[string]string{"url_code": code}).Decode(&result)
	if err!= nil {
        return "", err
    }
	err = mgr.redis_client.Set(context.TODO(), code, result.LongURL, 48*time.Hour).Err()
	if err!= nil {
        return "", err
    }
	return result.LongURL, err
}

func (mgr *manager) GetAndIncCounter() (int64, error) {
	count, err := mgr.protected_redis.Incr(context.TODO(), constant.COUNTER_KEY_REDIS).Result()
	if err != nil {
		return 0, err
	}
	return count, nil
}
