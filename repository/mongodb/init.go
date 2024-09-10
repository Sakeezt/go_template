package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Repository struct {
	Client  *mongo.Client
	DB      *mongo.Database
	Coll    *mongo.Collection
	URI     string
	DBName  string
	Timeout int
}

type Config struct {
	URI      string
	DBName   string
	CollName string
	Timeout  int
}

func New(ctx context.Context, config *Config) (repo *Repository, err error) {
	fullURI := fmt.Sprintf("%s/%s?authSource=admin", config.URI, config.DBName)
	clientOptions := options.Client().ApplyURI(fullURI)
	ctxTimeout, _ := context.WithTimeout(ctx, time.Duration(config.Timeout)*time.Second)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	if err = client.Connect(ctxTimeout); err != nil {
		return nil, err
	}

	if err = client.Ping(ctxTimeout, readpref.Primary()); err != nil {
		return nil, err
	}

	repo = &Repository{
		Client:  client,
		DB:      client.Database(config.DBName),
		Coll:    client.Database(config.DBName).Collection(config.CollName),
		URI:     config.URI,
		DBName:  config.DBName,
		Timeout: config.Timeout,
	}
	return repo, nil
}
