package mongod

import (
	"context"
	"fmt"
	"log"
	"time"

	"jimmie-go-lang/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ใช้กับตอน return กับตอนรับค่า
type Mongo interface {
	Delete(filter bson.M)
	Find(filter bson.M)
	FindOne(filter bson.D)
	UpdateOne(filter bson.M, updateValue bson.D)
	InsertOne(socialdata entity.Socialdata)
}

type SocialData struct {
	Host            string
	Port            string
	Username        string
	Password        string
	Database        string
	Collection      string
	mongoCtx        context.Context
	mongoCollection *mongo.Collection
}

func (mongoAdt SocialData) UpdateOne(filter bson.M, updateValue bson.D) {
	updateResult, err := mongoAdt.mongoCollection.UpdateOne(mongoAdt.mongoCtx, filter, updateValue)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

func (mongoAdt SocialData) FindOne(filter bson.D) entity.Socialdata {
	result := entity.Socialdata{}
	err := mongoAdt.mongoCollection.FindOne(mongoAdt.mongoCtx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func (mongoAdt SocialData) Find(filter bson.M) []entity.Socialdata {
	filterCursor, err := mongoAdt.mongoCollection.Find(mongoAdt.mongoCtx, filter)
	if err != nil {
		log.Fatal(err)
	}

	result := []entity.Socialdata{}
	if err = filterCursor.All(mongoAdt.mongoCtx, &result); err != nil {
		log.Fatal(err)
	}

	return result
}

func (mongoAdt SocialData) Delete(filter bson.M) {
	deleteResult, err := mongoAdt.mongoCollection.DeleteOne(mongoAdt.mongoCtx, filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}

func (mongoAdt *SocialData) Connection() {
	mongoUri := fmt.Sprintf("mongodb://%s:%s@%s:%s/admin",
		mongoAdt.Username,
		mongoAdt.Password,
		mongoAdt.Host,
		mongoAdt.Port,
	)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)

	db := client.Database(mongoAdt.Database)
	collection := db.Collection(mongoAdt.Collection)

	mongoAdt.mongoCollection = collection
}
