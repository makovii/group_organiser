package server

import (
	"github.com/gin-gonic/gin"
	"github.com/makovii/group_organiser/controller"
	database "github.com/makovii/group_organiser/db"
	"github.com/makovii/group_organiser/middleware"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	db := database.ConnectDatabase()

	user := controller.NewUserController(db)
	userGroup := router.Group("user")
	userGroup.Use(middleware.AuthMiddleware())
	userGroup.GET("/getUser", user.GetUser)
	userGroup.GET("/myNotifications", user.MyNotifications)
	userGroup.POST("/joinTeam", user.JoinTeam)
	userGroup.POST("/leaveTeam", user.LeaveTeam)
	userGroup.POST("/cancelRequest", user.CancelRequest)
	userGroup.GET("/getManagers", user.GetManagers)

	manager := controller.NewManagerController(db)
	managerGroup := router.Group("manager")
	managerGroup.Use(middleware.AuthMiddleware())
	managerGroup.POST("/createTeam", manager.CreateTeam)
	managerGroup.GET("/getAllteams", manager.GetAllTeams)
	managerGroup.GET("/getTeam", manager.GetTeam)
	managerGroup.PUT("/updateTeam", manager.UpdateTeam)
	managerGroup.DELETE("/deleteTeam", manager.DeleteTeam)

	return router
}
