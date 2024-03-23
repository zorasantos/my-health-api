package config

import "github.com/spf13/viper"

var (
	envVars *Environments
)

type Environments struct {
	DBSource  string `mapstructure:"DB_SOURCE"`
	SecretKey string `mapstructure:"SECRET_KEY"`
}

func LoadConfig(path string) (*Environments, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&envVars)
	if err != nil {
		panic(err)
	}

	return envVars, err
}

func GetEnvVars() *Environments {
	return envVars
}
