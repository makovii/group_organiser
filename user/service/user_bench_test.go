package service_test

import (
	"fmt"
	"net/http/httptest"
	"testing"

	// "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"github.com/makovii/group_organiser/config"
	"github.com/makovii/group_organiser/controller"
	"github.com/makovii/group_organiser/database"
	"github.com/makovii/group_organiser/middleware"
	userRepo "github.com/makovii/group_organiser/user/repository"
	"github.com/makovii/group_organiser/user/service"
)


func  BenchmarkUserGetUserById(b *testing.B) {

	cfg := config.GetConfig()
	db := database.ConnectDatabase()
	userRpository := userRepo.NewUserRepository(db, cfg)
	
	service := service.NewUserService(cfg, userRpository)

	for n := 0; n < b.N; n++ {
		if _, err := service.GetUserById(1); err != nil {
			fmt.Print(err)
		}
	}

}

func  BenchmarkUserGetNotifications(b *testing.B) {

	cfg := config.GetConfig()
	db := database.ConnectDatabase()
	userRpository := userRepo.NewUserRepository(db, cfg)
	
	service := service.NewUserService(cfg, userRpository)

	for n := 0; n < b.N; n++ {
		if _, err := service.GetNotifications(1); err != nil {
			fmt.Print(err)
		}
	}
	

}

func  BenchmarkUserGetAllManagers(b *testing.B) {

	cfg := config.GetConfig()
	db := database.ConnectDatabase()
	userRpository := userRepo.NewUserRepository(db, cfg)
	
	service := service.NewUserService(cfg, userRpository)

	for n := 0; n < b.N; n++ {
		if _, err := service.GetAllManagers(); err != nil {
			fmt.Print(err)
		}
	}

}

func  BenchmarkUserGetAllTeams(b *testing.B) {

	cfg := config.GetConfig()
	db := database.ConnectDatabase()
	userRpository := userRepo.NewUserRepository(db, cfg)
	
	service := service.NewUserService(cfg, userRpository)

	for n := 0; n < b.N; n++ {
		if _, err := service.GetAllTeams(); err != nil {
			fmt.Print(err)
		}
	}
}

func  BenchmarkUserJoinTeam(b *testing.B) {

	cfg := config.GetConfig()
	db := database.ConnectDatabase()
	userRpository := userRepo.NewUserRepository(db, cfg)
	
	service := service.NewUserService(cfg, userRpository)

	gin.SetMode(gin.TestMode)

	authedUser := middleware.AuthedUser{
		Id:             2,
		Name:           "John",
		Email:          "john@doe.com",
		Role: 					2,
	}

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Set("authedUser", authedUser)
	
	body := controller.BodyJoinTeam{ TeamId: 1 }
	for n := 0; n < b.N; n++ {
		if _, err := service.JoinTeam(c, body); err != nil {
			fmt.Print(err)
		}
	}

}