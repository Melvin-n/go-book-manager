package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	//bson is the format which mongodb data is stored in, this package is used to marshal/unmarshal between mongo and go understandable format
)

func Connect() {
	var mongoURI string = fmt.Sprintf("mongodb+srv://%v:%v@cluster0.mzjzk.mongodb.net/dummy?retryWrites=true&w=majority", os.Getenv("MONGO_USER"), os.Getenv("MONGO_PASSWORD"))

}

func init() {
	//load env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading env file")
	}
}
