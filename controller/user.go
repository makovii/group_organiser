package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/makovii/group_organiser/config"
	"github.com/makovii/group_organiser/database"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
	CFG *config.Config
}

type BodyJoinTeam struct {
	TeamId	uint	`json:"teamId"`
	UserId	uint	`json:"userId"`
}

func NewUserController(db *gorm.DB, cfg *config.Config) *UserController {
	return &UserController{DB: db, CFG: cfg}
}

func (u *UserController) GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	
	var user database.User
	u.DB.Where("id = ?", id).First(&user)

	c.JSON(http.StatusOK, user)
}

func (u *UserController) MyNotifications(c *gin.Context) {
	to, _ := strconv.Atoi(c.Query("to"))

	var requests []database.Request
	u.DB.Where("\"to\" = ?", to).Find(&requests)

	c.JSON(http.StatusOK, requests)
}

func (u *UserController) JoinTeam(c *gin.Context) {
	var body BodyJoinTeam
	
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var team database.Team
	u.DB.Where("id = ?", body.TeamId).First(&team)

	var request database.Request
	request.From = body.UserId
	request.To = team.ManagerID
	request.StatusId = uint(u.CFG.Status.WaitId)
	request.TypeId = uint(u.CFG.Type.LeaveTeamId)

	u.DB.Save(&request)
	c.JSON(http.StatusOK, request)
}

func (u *UserController) LeaveTeam(c *gin.Context) {
	var body BodyJoinTeam
	
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var team database.Team
	u.DB.Where("id = ?", body.TeamId).First(&team)

	var request database.Request
	request.From = body.UserId
	request.To = team.ManagerID
	request.StatusId = uint(u.CFG.Status.WaitId)
	request.TypeId = uint(u.CFG.Type.JoinTeamId)

	u.DB.Save(&request)
	c.JSON(http.StatusOK, request)
}

func (u *UserController) CancelRequest(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	var request database.Request
	u.DB.Where("id = ?", id).First(&request)

	request.StatusId = uint(u.CFG.Status.CancelId)
	u.DB.Save(&request)

	c.JSON(http.StatusOK, request)
}

func (u *UserController) GetManagers(c *gin.Context) {
	var managers []database.User

	u.DB.Where("role = ?", u.CFG.Role.ManagerId).Find(&managers)

	c.JSON(http.StatusOK, managers)
}
