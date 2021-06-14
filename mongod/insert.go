package mongod

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Insert() {
	type Socialdata struct {
		ID        primitive.ObjectID `bson:"_id,omitempty"`
		PageID    string             `bson:"page_id,omitempty"`
		Name      string             `bson:"name,omitempty"`
		CreatedAt time.Time          `bson:"created_at,omitempty"`
	}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/admin"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	db := client.Database("social-api")
	collection := db.Collection("data")

	socialdata := Socialdata{
		ID:        primitive.NewObjectID(),
		PageID:    "123456",
		Name:      "Jimmie page",
		CreatedAt: time.Now(),
	}

	insertResult, err := collection.InsertOne(ctx, socialdata)
	if err != nil {
		panic(err)
	}
	fmt.Println(insertResult.InsertedID)
}
