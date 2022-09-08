package mongo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	cli *mongo.Client
	db  *mongo.Database
)

// ConnectMongoDB function connect mongodb
func ConnectMongoDB(uri, database string) *mongo.Database {
	clientOptions := options.Client().ApplyURI(uri)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Ping right after connected
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	cli = client
	db = client.Database(database)
	return db
}

// GetDB func
func GetDB() *mongo.Database {
	return db
}

// Disconnect func
func Disconnect() {
	// Disconnect from MongoDB
	err := cli.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed!")
}
