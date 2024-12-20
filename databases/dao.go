package databases

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Aditya8840/Link/constant"
	"github.com/Aditya8840/Link/types"
)

func (mgr *manager) Insert(data *types.URL) error {
	instance := mgr.client.Database(constant.DATABASE).Collection(constant.COLLECTION_NAME)
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

	if err := mgr.evictIfNeeded(); err != nil {
        log.Printf("Warning: Cache eviction failed: %v", err)
    }
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

func (mgr *manager) evictIfNeeded() error {
	size, err := mgr.redis_client.DBSize(context.TODO()).Result()
	if err != nil {
		return err
	}

	if size >= mgr.maxCacheSize {
		numToRemove := int(float64(size) * 0.2)
		keys, err := mgr.redis_client.Keys(context.TODO(), "*").Result()
		if err != nil {
			return err
		}

		for i := 0; i < numToRemove && i < len(keys); i++ {
			if keys[i] == constant.COUNTER_KEY_REDIS {
				continue
			}
			_, err := mgr.redis_client.Del(context.TODO(), keys[i]).Result()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (mgr *manager) GetAndIncCounter() (int64, error) {
	count, err := mgr.redis_client.Incr(context.TODO(), constant.COUNTER_KEY_REDIS).Result()
	if err != nil {
		return 0, err
	}
	return count, nil
}
