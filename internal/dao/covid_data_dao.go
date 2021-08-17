package dao

import (
	"context"
	"covid19-india/internal/config"
	"covid19-india/internal/models"
	"errors"
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

func getCollection() *mgm.Collection {
	return mgm.CollectionByName(collectionId)
}

// PersistCovidData Bulk persist all covid data
func PersistCovidData(ctx context.Context, covidData []models.CovidData) error {
	collection := getCollection()

	var operations []mongo.WriteModel

	for _, data := range covidData {
		operation := mongo.NewUpdateOneModel()
		operation.SetFilter(bson.M{"region": data.Region})
		operation.SetUpdate(bson.M{"$set": data})
		operation.SetUpsert(true)

		operations = append(operations, operation)
	}

	if _, err := collection.BulkWrite(ctx, operations); err != nil {
		return err
	}

	return nil
}

// GetCovidDataForRegions Get covid data for multiple regions
func GetCovidDataForRegions(ctx context.Context, id []string) ([]models.CovidData, error) {
	collection := getCollection()

	var operations []bson.M

	for _, id := range id {
		if len(id) == 0 {
			continue
		}

		operation := bson.M{"region": id}
		operations = append(operations, operation)
	}

	if len(operations) == 0 {
		return nil, errors.New("no region found")
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

// GetCovidDataForRegion Get covid data for a region
func GetCovidDataForRegion(ctx context.Context, region string) (*models.CovidData, error) {
	collection := getCollection()

	if len(region) == 0 {
		return nil, nil
	}

	var data models.CovidData

	if err := collection.FindOne(ctx, bson.M{"region": region}).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
