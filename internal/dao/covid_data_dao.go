package dao

import (
	"context"
	"covid19-india/internal/config"
	"covid19-india/internal/models"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const dbId = "covid19"
const collectionId = "covid19-india"

func init() {
	err := mgm.SetDefaultConfig(
		&mgm.Config{CtxTimeout: 10 * time.Second},
		dbId,
		options.Client().ApplyURI(config.ENV.MongoUri),
	)

	if err != nil {
		log.Fatal(err)
	}
}

func getContext() context.Context {
	return mgm.Ctx()
}

func getCollection() *mgm.Collection {
	return mgm.CollectionByName(collectionId)
}

func PersistCovidData(covid3pDataset []models.Covid3pData) error {
	ctx := getContext()
	collection := getCollection()

	var operations []mongo.WriteModel

	for _, covid3pData := range covid3pDataset {
		covidData, err := covid3pData.ToCovidData()

		if err != nil {
			return err
		}

		operation := mongo.NewUpdateOneModel()
		operation.SetFilter(bson.M{"region": covidData.Region})
		operation.SetUpdate(bson.M{"$set": covidData})
		operation.SetUpsert(true)

		operations = append(operations, operation)
	}

	if _, err := collection.BulkWrite(ctx, operations); err != nil {
		return err
	}

	return nil
}

func GetCovidDataForStates(id []string) ([]models.CovidData, error) {
	ctx := getContext()
	collection := getCollection()

	var operations []bson.M

	for _, id := range id {
		operation := bson.M{"region": id}
		operations = append(operations, operation)
	}

	cursor, err := collection.Find(ctx, bson.M{"$or": operations})

	if err != nil {
		return nil, err
	}

	var data []models.CovidData

	if err := cursor.All(ctx, &data); err != nil {
		return nil, err
	}

	return data, nil
}
