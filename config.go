package main

import (
	"os"

	"github.com/silco-dev/vander/mongodb"
)

type Config struct {
	Token string
	Mongo mongodb.MongoConfig
	API   ApiConfig
}

func LoadConfig() *Config {
	mongo := mongodb.MongoConfig{
		ConnectionString: os.Getenv("MONGOURI"),
		Username:         os.Getenv("MONGOUSER"),
		Password:         os.Getenv("MONGOPASS"),
	}

	apiconfig := ApiConfig{
		Host:   os.Getenv("HOST"),
		Port:   os.Getenv("PORT"),
		Secret: os.Getenv("SECRET"),
	}

	return &Config{
		Token: os.Getenv("TOKEN"),
		Mongo: mongo,
		API:   apiconfig,
	}
}
