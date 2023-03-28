package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/makovii/group_organiser/database"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (u UserController) GetUser(c *gin.Context) {
	var user database.Player
	user.Id = 1
	user.Email = "someemail@gmail.com"
	user.Name = "Bob"
	user.Password = "42evev04vee9204vev02svvqvqev"
	c.JSON(http.StatusOK, user)
}

func (u UserController) MyNotifications(c *gin.Context) {
	var user database.Player
	//user.Notification = []string{"1", "2"}
	MyNotif := user.Notification
	c.JSON(http.StatusOK, MyNotif)
}

func (u UserController) JoinTeam(c *gin.Context) {
	var req database.Request
	req.Id = 1
	req.From = 24242
	req.StatusId = 1
	req.To = 222
	req.TypeId = 111
	c.JSON(http.StatusOK, req)
}

func (u UserController) LeaveTeam(c *gin.Context) {
	var req database.Request
	req.Id = 1
	req.From = 24242
	req.StatusId = 1
	req.To = 222
	req.TypeId = 222

	c.JSON(http.StatusOK, req)
}

func (u UserController) CancelRequest(c *gin.Context) {
	var req database.Request
	req.Id = 1
	req.From = 24242
	req.StatusId = 1
	req.To = 222
	req.TypeId = 333

	c.JSON(http.StatusOK, req)
}

func (u UserController) GetManagers(c *gin.Context) {
	var managers []database.Manager
	c.JSON(http.StatusOK, managers)
}
