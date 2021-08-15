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
	RedisUri       string
}

var ENV *AppConfig

func init() {
	ENV = &AppConfig{
		Port:           GetMandatoryEnvValue("PORT"),
		MongoUri:       GetMandatoryEnvValue("MONGODB_URI"),
		HereMapsApiKey: GetMandatoryEnvValue("HERE_MAPS_API_KEY"),
		RedisUri:       GetMandatoryEnvValue("REDIS_URI"),
	}
}

// GetMandatoryEnvValue Get Mandatory Environment value. Fatal if it is not present
func GetMandatoryEnvValue(key string) string {
	value := GetEnvValue(key)

	if len(value) == 0 {
		log.Fatal(fmt.Sprintf("Cannot find env variable %s", key))
	}

	return value
}

// GetEnvValue Get optional environment value
func GetEnvValue(key string) string {
	value := os.Getenv(key)

	return value
}
