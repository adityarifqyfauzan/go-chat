package config

import (
	"strings"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Params struct {
	DB  *gorm.DB
	Env *viper.Viper
}

// explicit call to Set > flag > env > config > key/value store > default
func New(env string, confFiles map[string]string) *viper.Viper {
	conf := viper.New()
	conf.SetDefault("environment", env)

	// Defaults: App
	conf.SetDefault("app.name", "default")
	conf.SetDefault("app.port.rest", "8080")
	conf.SetDefault("app.port.ws", "8082")

	// Defaults: DataBase
	conf.SetDefault("database.driver", "")
	conf.SetDefault("database.dsn", "") // If you use the MySQL driver with existing database client, you must create the client with parameter multiStatements=true:
	conf.SetDefault("database.auto_migrate", "off")

	// Conf Env
	conf.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "_", "__")) // APP_DATA__BASE_PASS -> app.data_base.pass
	conf.AutomaticEnv()                                              // Automatically load Env variables

	// Conf Files
	conf.SetConfigType("yaml") // We're using yaml
	conf.SetConfigName(env)    // Search for a config file that matches our environment
	conf.AddConfigPath("./")   // look for config in the working directory
	conf.ReadInConfig()        // Find and read the config file

	// Read additional files
	for confFile := range confFiles {
		conf.SetConfigName(confFile)
		conf.MergeInConfig()
	}

	return conf
}
