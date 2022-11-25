package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	DB     *mongo.Client
	Kanjis *mongo.Collection
	JMdict *mongo.Collection
}

func NewClient(uri, db string) (*Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	database := client.Database(db)
	kanjis := database.Collection("kanjis")
	jmdict := database.Collection("jmdict")

	return &Client{
		client,
		kanjis,
		jmdict,
	}, nil
}

func (client *Client) Close() error {
	return client.DB.Disconnect(context.TODO())
}
