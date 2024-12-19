package databases

import (
	"github.com/Aditya8840/Link/constant"
	"github.com/Aditya8840/Link/types"
)

func (mgr *manager) Insert(data interface{}, collectionName string) (interface{}, error) {
	instance := mgr.client.Database(constant.DATABASE).Collection(collectionName)
	result, err := instance.InsertOne(mgr.ctx, data)
	return result, err
}

func (mgr *manager) GetOriginalURL(code, collectionName string) (types.URL, error) {
	instance := mgr.client.Database(constant.DATABASE).Collection(collectionName)
    var url types.URL
    err := instance.FindOne(mgr.ctx, map[string]string{"url_code": code}).Decode(&url)
    return url, err
}