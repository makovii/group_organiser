package service_test

import (
	"testing"

	"github.com/makovii/group_organiser/database"
	"github.com/makovii/group_organiser/manager/repository"
	"github.com/makovii/group_organiser/manager/service"
)

func BenchmarkManagerCreateTeam(b *testing.B) {

	db := database.ConnectDatabase()
	managerRepository := repository.NewManagerRepository(db)

	service := service.NewManagerService(managerRepository)

	for n := 0; n < b.N; n++ {
		_, err := service.CreateTeam("Team A", 1)
		if err != nil {
			b.Errorf("failed to create team with error: %v", err)
		}
	}
}

func BenchmarkManagerGetAllTeams(b *testing.B) {

	db := database.ConnectDatabase()
	managerRepository := repository.NewManagerRepository(db)

	service := service.NewManagerService(managerRepository)

	for n := 0; n < b.N; n++ {
		_, err := service.GetAllTeams(1)
		if err != nil {
			b.Errorf("failed to get all teams with error: %v", err)
		}
	}
}

func BenchmarkManagerGetTeam(b *testing.B) {

	db := database.ConnectDatabase()
	managerRepository := repository.NewManagerRepository(db)

	service := service.NewManagerService(managerRepository)

	teamID := uint(1)
	managerID := uint(1)

	for n := 0; n < b.N; n++ {
		_, err := service.GetTeam(teamID, managerID)
		if err != nil {
			b.Errorf("failed to get team with error: %v", err)
		}
	}
}

func BenchmarkManagerUpdateTeam(b *testing.B) {

	db := database.ConnectDatabase()
	managerRepository := repository.NewManagerRepository(db)

	service := service.NewManagerService(managerRepository)

	teamID := uint(1)
	managerID := uint(1)
	newTeamName := "Updated Team A"

	for n := 0; n < b.N; n++ {
		_, err := service.UpdateTeam(teamID, managerID, newTeamName)
		if err != nil {
			b.Errorf("failed to update team with error: %v", err)
		}
	}
}

func BenchmarkManagerDeleteTeam(b *testing.B) {

	db := database.ConnectDatabase()
	managerRepository := repository.NewManagerRepository(db)

	service := service.NewManagerService(managerRepository)

	teamID := uint(1)
	managerID := uint(1)

	for n := 0; n < b.N; n++ {
		err := service.DeleteTeam(teamID, managerID)
		if err != nil {
			b.Errorf("failed to delete team with error: %v", err)
		}
	}
}
