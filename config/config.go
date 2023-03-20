package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server   *ServerConfig
	Secrets  *SecretConfig
	Database *DatabaseConfig
}

type ServerConfig struct {
	Secure bool
	Domain string
	Host   string
	Port   int64
}

type SecretConfig struct {
	Secret string
}

type DatabaseConfig struct {
	Host     string
	Port     int64
	User     string
	Password string
	DbName   string
	Secure   bool
}

var config *Config

func InitConfig(name string) *Config {
	viper.AddConfigPath("config")
	viper.SetConfigName(name)
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Error while reading config: %s", err.Error()))
	}

	serverConfig := ServerConfig{
		Secure: viper.Get("server.secure").(bool),
		Domain: viper.Get("server.domain").(string),
		Host:   viper.Get("server.host").(string),
		Port:   viper.Get("server.port").(int64),
	}

	secretConfig := SecretConfig{
		Secret: viper.Get("secrets.jwt").(string),
	}

	dbConfig := DatabaseConfig{
		Host:     viper.Get("db.host").(string),
		Port:     viper.Get("db.port").(int64),
		User:     viper.Get("db.user").(string),
		Password: viper.Get("db.password").(string),
		DbName:   viper.Get("db.name").(string),
		Secure:   viper.Get("db.secure").(bool),
	}

	config = &Config{
		Server:   &serverConfig,
		Secrets:  &secretConfig,
		Database: &dbConfig,
	}

	return config
}

func GetConfig() *Config {
	if config != nil {
		return config
	}
	return InitConfig("config")
}
