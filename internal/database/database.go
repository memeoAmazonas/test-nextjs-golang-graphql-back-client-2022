package database

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
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
func GetConnection() error {
	url_db := fmt.Sprintf("%s:%s", os.Getenv("URL_DB"), os.Getenv("PORT_DB"))
	client, ctx, cancel, err := Connect(url_db)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	defer Close(client, ctx, cancel)
	return ping(client, ctx)
}
func SaveOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	collection := client.Database(dataBase).Collection(col)
	return collection.InsertOne(ctx, doc)
}
func Query(client *mongo.Client, ctx context.Context, dataBase, col string, filter bson.D, order *options.FindOptions) (result *mongo.Cursor, err error) {
	collection := client.Database(dataBase).Collection(col)
	result, err = collection.Find(ctx, filter, order)
	return result, err
}
func FindOne(client *mongo.Client, ctx context.Context, dataBase, col string, filter bson.D) (result *mongo.SingleResult, err error) {
	collection := client.Database(dataBase).Collection(col)
	one := collection.FindOne(ctx, filter)
	return one, nil
}
func UpdateOne(client *mongo.Client, ctx context.Context, dataBase, col string, condition, value bson.D) (*mongo.UpdateResult, error) {
	collection := client.Database(dataBase).Collection(col)
	one, err := collection.UpdateOne(ctx, condition, value)
	if err != nil {
		return nil, err
	}
	return one, nil
}
func DeleteOne(client *mongo.Client, ctx context.Context, dataBase, col string, condition bson.D) (*mongo.DeleteResult, error) {
	collection := client.Database(dataBase).Collection(col)
	one, err := collection.DeleteOne(ctx, condition)
	if err != nil {
		return nil, err
	}
	return one, nil
}
