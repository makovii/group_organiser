package server

import (

	"github.com/gin-gonic/gin"
	"github.com/makovii/group_organiser/middleware"
	"github.com/makovii/group_organiser/controller"
)


func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	
	user := cotroller.NewUserController();

	userGroup := router.Group("user")
	userGroup.Use(middleware.AuthMiddleware())
	userGroup.GET("/getUser", user.GetUser)
	userGroup.GET("/myNotifications", user.MyNotifications)
	userGroup.POST("/joinTeam", user.JoinTeam)
	userGroup.POST("/leaveTeam", user.LeaveTeam)
	userGroup.POST("/cancelRequest", user.CancelRequest)
	userGroup.GET("/getManagers", user.GetManagers)

	return router
}