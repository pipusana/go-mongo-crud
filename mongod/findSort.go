package mongod

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindSort() {
	type Socialdata []struct {
		ID     primitive.ObjectID `bson:"_id, omitempty"`
		Pageid string             `bson:"page_id, omitempty"`
		Name   string             `bson:"name, omitempty"`
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

	opts := options.Find()
	opts.SetSort(bson.D{{"created_at", 1}})

	sortCursor, err := collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		log.Fatal(err)
	}
	result := Socialdata{}
	if err = sortCursor.All(ctx, &result); err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
