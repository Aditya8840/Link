package databases

import (
	"time"

	"github.com/Aditya8840/Link/constant"
	"github.com/Aditya8840/Link/types"
)

func (mgr *manager) Insert(data *types.URL) error {
	instance := mgr.client.Database(constant.DATABASE).Collection(constant.COLLECTION_NAME)
	_, err := instance.InsertOne(mgr.ctx, data)
	if err != nil {
        return err
    }

	err = mgr.redis_client.Set(
		mgr.ctx,
		data.URLCode,
		data.LongURL,
		48*time.Hour,
	).Err()
	return err
}

func (mgr *manager) GetOriginalURL(code string) (string, error) {
	longUrl, err := mgr.redis_client.Get(mgr.ctx, code).Result()
	if err == nil {
		return longUrl, nil
	}
	instance := mgr.client.Database(constant.DATABASE).Collection(constant.COLLECTION_NAME)
    var url types.URL
    err = instance.FindOne(mgr.ctx, map[string]string{"url_code": code}).Decode(&url)
    return url.LongURL, err
}

func (mgr *manager) evictIfNeeded() error {
	size, err := mgr.redis_client.DBSize(mgr.ctx).Result()
    if err != nil {
        return err
    }

	if size >= mgr.maxCacheSize {
		numToRemove := int(float64(size) * 0.2)

		keys, err := mgr.redis_client.Keys(mgr.ctx, "*").Result()
		if err != nil {
			return err
		}

		for i := 0; i < numToRemove && i < len(keys); i++ {
			mgr.redis_client.Del(mgr.ctx, keys[i])
		}
	}
}