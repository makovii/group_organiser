package server

import (
	"github.com/gin-gonic/gin"
	"github.com/makovii/group_organiser/controller"
	database "github.com/makovii/group_organiser/db"
	"github.com/makovii/group_organiser/middleware"
	"github.com/makovii/group_organiser/config"
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


	user := controller.NewUserController(db)
	userGroup := router.Group("user")
	userGroup.Use(middleware.IsAuthorized(cfg))
	userGroup.GET("/getUser", user.GetUser)
	userGroup.GET("/myNotifications", user.MyNotifications)
	userGroup.POST("/joinTeam", user.JoinTeam)
	userGroup.POST("/leaveTeam", user.LeaveTeam)
	userGroup.POST("/cancelRequest", user.CancelRequest)
	userGroup.GET("/getManagers", user.GetManagers)

	manager := controller.NewManagerController(db)
	managerGroup := router.Group("manager")
	managerGroup.Use(middleware.IsAuthorized(cfg))
	managerGroup.POST("/createTeam", manager.CreateTeam)
	managerGroup.GET("/getAllteams", manager.GetAllTeams)
	managerGroup.GET("/getTeam", manager.GetTeam)
	managerGroup.PUT("/updateTeam", manager.UpdateTeam)
	managerGroup.DELETE("/deleteTeam", manager.DeleteTeam)

	admin := controller.NewAdminController(db)
	adminGroup := router.Group("admin")
	adminGroup.Use(middleware.IsAuthorized(cfg))
	adminGroup.GET("/getAdmin", admin.GetAdmin)
	adminGroup.GET("/getAdminById", admin.GetById)
	adminGroup.POST("/banById", admin.BanById)
	adminGroup.POST("/getTeams", admin.GetTeams)

	return router
}
