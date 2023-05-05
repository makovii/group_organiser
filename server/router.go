package server

import (
	"github.com/gin-gonic/gin"
	"github.com/makovii/group_organiser/config"
	"github.com/makovii/group_organiser/controller"
	"github.com/makovii/group_organiser/database"
	"github.com/makovii/group_organiser/middleware"
	"github.com/makovii/group_organiser/repository"
	"github.com/makovii/group_organiser/service"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	cfg := config.GetConfig()
	db := database.ConnectDatabase()

	auth := controller.NewAuthController(db, cfg)
	authGroup := router.Group("auth")
	authGroup.POST("/signIn", auth.SignIn)
	authGroup.POST("/signUp", auth.SignUp)

	userRpository := repository.NewUserRepository(db, cfg)
	userService := service.NewUserService(userRpository)
	user := controller.NewUserController(db, cfg, userService)
	userGroup := router.Group("user")
	userGroup.Use(middleware.IsAuthorized(cfg))
	userGroup.GET("/getUser", user.GetUserById)
	userGroup.GET("/myNotifications", user.GetNotifications)
	userGroup.POST("/joinTeam", user.JoinTeam)
	userGroup.POST("/leaveTeam", user.LeaveTeam)
	userGroup.POST("/cancelRequest", user.CancelRequest)
	userGroup.GET("/getAllManagers", user.GetAllManagers)
	userGroup.GET("/getAllTeams", user.GetAllTeams)

	managerRepository := repository.NewManagerRepository(db)
	managerService := service.NewManagerService(managerRepository)
	manager := controller.NewManagerController(db, cfg, managerService)
	managerGroup := router.Group("manager")
	managerGroup.Use(middleware.IsAuthorized(cfg))
	managerGroup.POST("/createTeam", manager.CreateTeam)
	managerGroup.GET("/getAllteams", manager.GetAllTeams)
	managerGroup.GET("/getTeam", manager.GetTeam)
	managerGroup.PUT("/updateTeam", manager.UpdateTeam)
	managerGroup.DELETE("/deleteTeam", manager.DeleteTeam)

	admin := controller.NewAdminController(db, cfg)
	adminGroup := router.Group("admin")
	adminGroup.Use(middleware.IsAuthorized(cfg))
	adminGroup.GET("/getAdmins", admin.GetAdmins)
	adminGroup.GET("/getUserById", admin.GetUserById)
	adminGroup.GET("/getTeamById", admin.GetTeamById)
	adminGroup.POST("/banById", admin.BanById)
	adminGroup.POST("/getTeams", admin.GetTeams)
	adminGroup.POST("/acceptRegistration", admin.AcceptManagerRegistration)

	return router
}
