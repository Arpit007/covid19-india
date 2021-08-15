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
	RedisPassword  string
}

var ENV *AppConfig

func init() {
	ENV = &AppConfig{
		Port:           getReqdEnvValue("PORT"),
		MongoUri:       getReqdEnvValue("MONGODB_URI"),
		HereMapsApiKey: getReqdEnvValue("HERE_MAPS_API_KEY"),
		RedisUri:       getReqdEnvValue("REDIS_URI"),
		RedisPassword:  getEnvValue("REDIS_PASSWORD"),
	}
}

func getReqdEnvValue(key string) string {
	value := getEnvValue(key)

	if len(value) == 0 {
		log.Fatal(fmt.Sprintf("Cannot find env variable %s", key))
	}

	return value
}

func getEnvValue(key string) string {
	value := os.Getenv(key)

	return value
}
