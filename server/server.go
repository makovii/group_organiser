package server

import (
	"fmt"

	"github.com/makovii/group_organiser/config"
)

func Init() {
	router := NewRouter()

	cfg := config.GetConfig()

	router.Run(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
}
