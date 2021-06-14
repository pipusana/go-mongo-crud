package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Socialdata struct {
	ID        primitive.ObjectID `bson:"_id, omitempty"`
	PageID    string             `bson:"page_id, omitempty"`
	Name      string             `bson:"name, omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
}
