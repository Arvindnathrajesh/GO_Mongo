package domain

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

func ConnDB() {
	uri := "mongodb+srv://ArvindNath:pass%40123@cluster0.9tist.mongodb.net/?retryWrites=true&w=majority"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	db = client.Database("goMongo")
	fmt.Println("Successfully connected to MongoDB")
}
