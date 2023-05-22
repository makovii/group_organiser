package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Admin		 *AdminConfig
	Status	 *StatusConfig
	Type		 *TypeConfig
	Role		 *RoleConfig
	Server   *ServerConfig
	Secrets  *SecretConfig
	Database *DatabaseConfig
}

type AdminConfig struct {
	Id	int64
}

type StatusConfig struct {
	WaitId int64
	AcceptId int64
	RejectId int64
	CancelId int64
}

type TypeConfig struct {
	RegistrationId	int64
	JoinTeamId	int64
	LeaveTeamId	int64
}

type RoleConfig struct {
	AdminId int64
	ManagerId int64
	PlayerId int64
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

	adminConfig := AdminConfig{
		Id: viper.Get("admin.admin_id").(int64),
	}

	statusConfig := StatusConfig{
		WaitId: viper.Get("status.waitId").(int64),
		AcceptId: viper.Get("status.acceptId").(int64),
		RejectId: viper.Get("status.rejectId").(int64),
		CancelId: viper.Get("status.CancelId").(int64),
	}

	typeConfig := TypeConfig{
		RegistrationId: viper.Get("type.registrationId").(int64),
		JoinTeamId:	viper.Get("type.JoinTeamId").(int64),
		LeaveTeamId: viper.Get("type.LeaveTeamId").(int64),
	}

	roleConfig := RoleConfig{
		AdminId: viper.Get("role.adminId").(int64),
		ManagerId: viper.Get("role.managerId").(int64),
		PlayerId: viper.Get("role.playerId").(int64),
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
		Admin:		&adminConfig,
		Status: 	&statusConfig,
		Type:			&typeConfig,
		Role:			&roleConfig,
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
