package handle

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func EnsureCollectionExists(collection *mongo.Collection) error {
	names, err := collection.Database().ListCollectionNames(context.Background(), bson.M{"name": collection.Name()})
	if err != nil {
		return err
	}

	if len(names) == 0 {
		err = collection.Database().CreateCollection(context.Background(), collection.Name())
		if err != nil {
			return err
		}
	}

	return nil
}
