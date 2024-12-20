package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type URL struct {
	ID primitive.ObjectID `bson:"_id, omitempty"`
	URLCode string `bson:"url_code"`
	LongURL string `bson:"long_url"`
	CreatedAt int64 `bson:"created_at"`
}