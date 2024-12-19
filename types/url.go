package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type URL struct {
	ID primitive.ObjectID `bson:"_id, omitempty"`
	UrlCode string `bson:"url_code"`
	LongURL string `bson:"long_url"`
	CreatedAt int64 `bson:"created_at"`
	ExpiresAt int64 `bson:"expires_at"`
}