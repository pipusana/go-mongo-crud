package main

import (
	"fmt"
	"jimmie-go-lang/mongod"
	"time"

	"jimmie-go-lang/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	host := "localhost"
	port := "27017"
	username := "root"
	password := "root"
	database := "social-api"
	collection := "data"

	SocialData := mongod.SocialData{
		Host:       host,
		Port:       port,
		Username:   username,
		Password:   password,
		Database:   database,
		Collection: collection,
	}
	SocialData.Connection()

	// Insert One
	socialdata := entity.Socialdata{
		ID:        primitive.NewObjectID(),
		PageID:    "123456",
		Name:      "Jimmie page",
		CreatedAt: time.Now(),
	}
	insertedID := SocialData.InsertOne(socialdata)

	// Find All
	result := SocialData.Find(bson.M{})
	fmt.Printf("%v\n", result)

	// Update One
	filterUpdateOne := bson.M{"_id": insertedID}
	updateValue := bson.D{
		{"$set", bson.D{
			{"page_id", "191"},
		}},
	}
	SocialData.UpdateOne(filterUpdateOne, updateValue)

	// Find One
	filterFindOne := bson.D{
		{"page_id", "191"},
		{"name", "Jimmie page"},
	}
	resultValue := SocialData.FindOne(filterFindOne)
	fmt.Printf("%v\n", resultValue)

	// Delete
	filter := bson.M{"_id": insertedID}
	SocialData.Delete(filter)

	// Find All
	result = SocialData.Find(bson.M{})
	fmt.Printf("%v\n", result)

}
