package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

type AppConfig struct {
	Port           string
	MongoUri       string
	HereMapsApiKey string
}

var ENV *AppConfig

func init() {
	ENV = &AppConfig{
		Port:           GetEnvValue("PORT"),
		MongoUri:       GetEnvValue("MONGODB_URI"),
		HereMapsApiKey: GetEnvValue("HERE_MAPS_API_KEY"),
	}
}

func GetEnvValue(key string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		log.Fatal(fmt.Sprintf("Cannot find env variable %s", key))
	}

	return value
}
