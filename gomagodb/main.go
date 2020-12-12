package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	mongoPassword := os.Getenv("MANGODB_PASSWORD")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mangodbuser:" + mongoPassword + "@cluster0.ezd6m.mongodb.net/<dbname>?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	restaurantsDB := client.Database("sample_restaurants")
	restaurantsCollection := restaurantsDB.Collection("restaurants")

	objID, _ := primitive.ObjectIDFromHex("5eb3d669b31de5d588f48c35")
	cur, err := restaurantsCollection.Find(ctx, bson.M{"_id": objID})
	if err != nil {
		log.Fatal(err)
	}

	var restaurants []bson.M
	if err = cur.All(ctx, &restaurants); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(restaurants)
	s, _ := json.MarshalIndent(restaurants, "", "\t")
	fmt.Print(string(s))

}
