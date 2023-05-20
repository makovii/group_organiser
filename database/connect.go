package database

import (
	"fmt"

	"github.com/makovii/group_organiser/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	cfg := config.GetConfig()

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	cfg.Database.Host, cfg.Database.Port, cfg.Database.User,
	// 	cfg.Database.Password, cfg.Database.DbName)
	psqlInfo := fmt.Sprintf(cfg.Database.Dbstring)

	database, err := gorm.Open(postgres.Open(psqlInfo))
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database

	err = database.AutoMigrate(&Type{})
	if err != nil {
		return nil
	}

	err = database.AutoMigrate(&Status{})
	if err != nil {
		return nil
	}

	err = database.AutoMigrate(&Request{})
	if err != nil {
		return nil
	}

	err = database.AutoMigrate(&Team{})
	if err != nil {
		return nil
	}

	err = database.AutoMigrate(&User{})
	if err != nil {
		return nil
	}

	err = database.AutoMigrate(&Role{})
	if err != nil {
		return nil
	}

	return database
}
