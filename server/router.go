package server

import (
	"github.com/gin-gonic/gin"
	"github.com/makovii/group_organiser/config"
	"github.com/makovii/group_organiser/controller"
	"github.com/makovii/group_organiser/database"
	"github.com/makovii/group_organiser/middleware"
	adminService "github.com/makovii/group_organiser/admin/service"
	managerService "github.com/makovii/group_organiser/manager/service"
	userService "github.com/makovii/group_organiser/user/service"
	adminRepo "github.com/makovii/group_organiser/admin/repository"
	managerRepo "github.com/makovii/group_organiser/manager/repository"
	userRepo "github.com/makovii/group_organiser/user/repository"
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

	userRpository := userRepo.NewUserRepository(db, cfg)
	userService := userService.NewUserService(cfg, userRpository)
	user := controller.NewUserController(userService)
	userGroup := router.Group("user")
	userGroup.Use(middleware.IsAuthorized(cfg))
	userGroup.GET("/getUser", user.GetUserById)
	userGroup.GET("/myNotifications", user.GetNotifications)
	userGroup.POST("/joinTeam", user.JoinTeam)
	userGroup.POST("/leaveTeam", user.LeaveTeam)
	userGroup.POST("/cancelRequest", user.CancelRequest)
	userGroup.GET("/getAllManagers", user.GetAllManagers)
	userGroup.GET("/getAllTeams", user.GetAllTeams)

	managerRepository := managerRepo.NewManagerRepository(db)
	managerService := managerService.NewManagerService(managerRepository)
	manager := controller.NewManagerController(db, cfg, managerService)
	managerGroup := router.Group("manager")
	managerGroup.Use(middleware.IsAuthorized(cfg))
	managerGroup.POST("/createTeam", manager.CreateTeam)
	managerGroup.GET("/getAllteams", manager.GetAllTeams)
	managerGroup.GET("/getTeam", manager.GetTeam)
	managerGroup.PUT("/updateTeam", manager.UpdateTeam)
	managerGroup.DELETE("/deleteTeam", manager.DeleteTeam)

	adminRepository := adminRepo.NewAdminRepository(db, cfg)
	adminService := adminService.NewAdminService(adminRepository)
	admin := controller.NewAdminController(db, cfg, adminService)
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
