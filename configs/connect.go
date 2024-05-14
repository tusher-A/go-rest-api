package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB() *mongo.Client {
	// create context to wait for 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(EnvDatabaseURI()).SetServerAPIOptions(serverAPI);
	client, err := mongo.Connect(ctx, clientOptions)
	if err!=nil {
		log.Fatal(err)
	}
	// defer disconnect so that connection gets closed
	defer func(){
		if err:=client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err:= client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to db")
	return client;
}

// client instance
var DB = ConnectDB();

func getCollection(client *mongo.Client, collectionName string) *mongo.Collection  {
	collection := client.Database("GoRestAPI").Collection(collectionName)
	return collection
}