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

func FindOne() {
	type Socialdata struct {
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

	// id, _ := primitive.ObjectIDFromHex("60bb2c4a772f6eda9f10e337")
	// filter := bson.M{"_id": id}

	// filter := bson.M{
	// 	"page_id": "1",
	// 	"name":    "Jimmie page",
	// }

	filter := bson.D{
		{"page_id", "1"},
		{"name", "Jimmie page"},
	}

	result := Socialdata{}
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("%s", result.ID)
	fmt.Println("%v", result.Name)
	fmt.Println("%v", result.Pageid)
}

//  Note: bson.D vs bson.M https://stackoverflow.com/questions/64281675/bson-d-vs-bson-m-for-find-queries
