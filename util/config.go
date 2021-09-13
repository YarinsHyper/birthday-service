package util

import (
	"github.com/spf13/viper"
)

const (
	// ConfigMongoConnectionString is ..
	ConfigMongoConnectionString = "mongo_host"
	// configHealthCheckInterval is ..
	configHealthCheckInterval = "health_check_interval"
)

// Config is used to declare every variable in
// app.env file in order to use it in go files
type Config struct {
	MongoURL string `mapstructure:"MONGO_URL"`
}

// LoadConfig loads all variables from app.env file
func LoadConfig() (err error) {
	viper.SetDefault(ConfigMongoConnectionString, "mongodb://root:example@0.0.0.0:27017")
	viper.SetDefault(configHealthCheckInterval, 3)
	viper.AutomaticEnv()

	return nil
}
