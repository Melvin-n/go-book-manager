package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	//bson is the format which mongodb data is stored in, this package is used to marshal/unmarshal between mongo and go understandable format
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var db *mongo.Database

func Connect() {
	var mongoURI string = fmt.Sprintf("mongodb+srv://%v:%v@cluster0.mzjzk.mongodb.net/dummy?retryWrites=true&w=majority", os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"))
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	//timeout for connecting
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(ctx)

	//pings the mongodb cluster, if no error is return then connection was succesful
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(databases)

	db = client.Database("book-collection")

}

func GetDB() *mongo.Database {
	return db
}

func init() {
	//load env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading env file")
	}
}
