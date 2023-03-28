package main

import (
	"github.com/makovii/group_organiser/server"
	"fmt"
	"github.com/makovii/group_organiser/config"
)

func main() {
	router := server.NewRouter()

	cfg := config.GetConfig()

	err := router.Run(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
	if err != nil {
		fmt.Println("Smth wrong witn router.Run function: ", err)
	}
}
