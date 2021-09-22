package util

import (
	"github.com/spf13/viper"
)

const (
	// ConfigMongoConnectionString is ..
	ConfigMongoConnectionString = "mongo_host"
)

// Config is used to declare every variable in
// app.env file in order to use it in go files
type Config struct {
	MongoURL string `mapstructure:"MONGO_URL"`
}

// LoadConfig loads all variables from app.env file
func LoadConfig() (err error) {
	viper.SetDefault(ConfigMongoConnectionString, "mongodb://168.63.40.248:27017/birthdayDB")
	viper.AutomaticEnv()

	return nil
}
