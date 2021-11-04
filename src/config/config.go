package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configurations exported
type Configurations struct {
	Server       ServerConfigurations
	Database     DatabaseConfigurations
	Discord      DiscordConfigurations
	EXAMPLE_PATH string
	EXAMPLE_VAR  string
}

// ServerConfigurations exported
type ServerConfigurations struct {
	Port int
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
}

type DiscordConfigurations struct {
	Prefix string
	Token  string
}

func InitConfig(configname string) Configurations {
	// Set the file name of the configurations file
	viper.SetConfigName(configname)

	// Set the path to look for the configurations file
	viper.AddConfigPath(".//src//config")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("database.dbname", "test_db")

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return configuration
}
