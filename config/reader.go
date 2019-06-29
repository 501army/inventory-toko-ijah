package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// ConfigReader is
type ConfigReader struct{}

//Read is
func (w *ConfigReader) Read() {
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.AddConfigPath("./config") // path to look for the config file in
	viper.SetConfigType("json")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}
