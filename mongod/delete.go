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

func Delete() {
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

	id, _ := primitive.ObjectIDFromHex("60bb39d3c73dd73c5007039a")
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}
