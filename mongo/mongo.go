package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Connection struct {
	Host     string
	DbName   string
	User     string
	Password string
	Port     string
}

type Client struct {
	MongoDB *mongo.Database
	MongoClient *mongo.Client
}

func (c Connection) DBConnect() (client *mongo.Client, err error) {
	atlasURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", c.User, c.Password, c.Host, c.Port)
	client, err = mongo.NewClient(options.Client().ApplyURI(atlasURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())

	return client, err
}

func(c Connection) DBDisconnect(client *mongo.Client){
	client.Disconnect(context.TODO())
}
