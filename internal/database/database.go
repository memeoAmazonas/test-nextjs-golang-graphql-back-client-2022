package database

import (
	"context"
	log "github.com/sirupsen/logrus"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error(err.Error())
		}
	}()
}
func Connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}
func ping(client *mongo.Client, ctx context.Context) error {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	log.Info("database connected successfully")
	return nil
}
func GetConnection() {
	client, ctx, cancel, err := Connect("mongodb://localhost:27017")
	if err != nil {
		log.Error(err.Error())
	}
	defer Close(client, ctx, cancel)
	ping(client, ctx)
}
func SaveOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	collection := client.Database(dataBase).Collection(col)
	return collection.InsertOne(ctx, doc)
}
func Query(client *mongo.Client, ctx context.Context, dataBase, col string, filter bson.D) (result *mongo.Cursor, err error) {
	collection := client.Database(dataBase).Collection(col)
	result, err = collection.Find(ctx, filter)
	return result, err
}
