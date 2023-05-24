package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/makovii/group_organiser/config"
	"github.com/makovii/group_organiser/server"
)

func main() {
	router := server.NewRouter()

	cfg := config.GetConfig()

	portStr := os.Getenv("PORT")
	portInt64 := cfg.Server.Port
	if portStr != "" {
		var err error
			portInt64, err = strconv.ParseInt(portStr, 10, 64) // Default port if not specified
			if err != nil {
				fmt.Println("Error during conversion")
				return
			}
	}

	err := router.Run(fmt.Sprintf("%s:%d", cfg.Server.Host, portInt64))
	if err != nil {
		fmt.Println("Smth wrong witn router.Run function: ", err)
	}
}
