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

func InsertMany() {
	type mongoInsertMany []interface {
	}

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

	insertValue1 := Socialdata{
		ID:        primitive.NewObjectID(),
		PageID:    "123456",
		Name:      "Jimmie page",
		CreatedAt: time.Now(),
	}

	insertValue2 := Socialdata{
		ID:        primitive.NewObjectID(),
		PageID:    "567890",
		Name:      "net page",
		CreatedAt: time.Now(),
	}

	insertValue := mongoInsertMany{insertValue1, insertValue2}
	insertResult, err := collection.InsertMany(ctx, insertValue)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted %v documents into episode collection!\n", len(insertResult.InsertedIDs))
}
