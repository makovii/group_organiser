package service_test

import (
	"fmt"
	"testing"

	"github.com/makovii/group_organiser/admin/repository"
	"github.com/makovii/group_organiser/admin/service"
	"github.com/makovii/group_organiser/config"
	"github.com/makovii/group_organiser/database"
)

func BenchmarkUserGetAdmins(b *testing.B) {

	cfg := config.GetConfig()
	db := database.ConnectDatabase()
	adminRepository := repository.NewAdminRepository(db, cfg)

	service := service.NewAdminService(adminRepository)

	for n := 0; n < b.N; n++ {
		if _, err := service.GetAdmins(); err != nil {
			fmt.Print(err)
		}
	}

}

func BenchmarkGetUserById(b *testing.B) {

	cfg := config.GetConfig()
	db := database.ConnectDatabase()
	adminRepository := repository.NewAdminRepository(db, cfg)

	service := service.NewAdminService(adminRepository)

	for n := 0; n < b.N; n++ {
		if _, err := service.GetUserById(1); err != nil {
			fmt.Print(err)
		}
	}

}

func BenchmarkUserGetTeamById(b *testing.B) {

	cfg := config.GetConfig()
	db := database.ConnectDatabase()
	adminRepository := repository.NewAdminRepository(db, cfg)

	service := service.NewAdminService(adminRepository)

	for n := 0; n < b.N; n++ {
		if _, err := service.GetTeamById(1); err != nil {
			fmt.Print(err)
		}
	}
}

func BenchmarkUserGetTeams(b *testing.B) {

	cfg := config.GetConfig()
	db := database.ConnectDatabase()
	adminRepository := repository.NewAdminRepository(db, cfg)

	service := service.NewAdminService(adminRepository)

	for n := 0; n < b.N; n++ {
		if _, err := service.GetTeams(); err != nil {
			fmt.Print(err)
		}
	}
}

func BenchmarkBanById(b *testing.B) {

	cfg := config.GetConfig()
	db := database.ConnectDatabase()
	adminRepository := repository.NewAdminRepository(db, cfg)

	service := service.NewAdminService(adminRepository)

	for n := 0; n < b.N; n++ {
		if _, err := service.BanById(1); err != nil {
			fmt.Print(err)
		}
	}
}

func BenchmarkAcceptManagerRegistration(b *testing.B) {

	cfg := config.GetConfig()
	db := database.ConnectDatabase()
	adminRepository := repository.NewAdminRepository(db, cfg)

	service := service.NewAdminService(adminRepository)

	for n := 0; n < b.N; n++ {
		if _, err := service.AcceptManagerRegistration(1); err != nil {
			fmt.Print(err)
		}
	}
}
