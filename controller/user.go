package cotroller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{ DB: db }
}

func (u UserController) GetUser(c *gin.Context) {
	user := "user1"
	c.JSON(http.StatusOK, user)
	return
}

func (u UserController) MyNotifications(c *gin.Context) {
	MyNotifications := "MyNotifications"
	c.JSON(http.StatusOK, MyNotifications)
	return
}

func (u UserController) JoinTeam(c *gin.Context) {
	JoinTeam := "useJoinTeamr1"
	c.JSON(http.StatusOK, JoinTeam)
	return
}

func (u UserController) LeaveTeam(c *gin.Context) {
	LeaveTeam := "LeaveTeam"
	c.JSON(http.StatusOK, LeaveTeam)
	return
}

func (u UserController) CancelRequest(c *gin.Context) {
	CancelRequest := "CancelRequest"
	c.JSON(http.StatusOK, CancelRequest)
	return
}

func (u UserController) GetManagers(c *gin.Context) {
	GetManagers := "GetManagers"
	c.JSON(http.StatusOK, GetManagers)
	return
}