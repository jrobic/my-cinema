package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBMongo struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewDBMongo(uri, dbName string, timeout time.Duration) (*DBMongo, error) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)

	return &DBMongo{
		Client: client,
		DB:     db,
	}, nil
}

func (db *DBMongo) Close() {
	if err := db.Client.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}
