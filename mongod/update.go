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

func Update() {
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

	id, _ := primitive.ObjectIDFromHex("60bb2c4a772f6eda9f10e337")
	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{
			{"page_id", "1"},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}
