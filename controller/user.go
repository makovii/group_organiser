package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/makovii/group_organiser/config"
	"github.com/makovii/group_organiser/database"
	"github.com/makovii/group_organiser/middleware"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
	CFG *config.Config
}

type BodyJoinTeam struct {
	TeamId	uint	`json:"teamId"`
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

	authedUser, _ := c.Get("authedUser")
	user := authedUser.(middleware.AuthedUser)

	var team database.Team
	u.DB.Where("id = ?", body.TeamId).First(&team)

	if team.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	var request database.Request
	request.From = uint(user.Id)
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

	authedUser, _ := c.Get("authedUser")
	user := authedUser.(middleware.AuthedUser)

	var team database.Team
	u.DB.Where("id = ?", body.TeamId).First(&team)

	if team.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	var request database.Request
	request.From = uint(user.Id)
	request.To = team.ManagerID
	request.StatusId = uint(u.CFG.Status.WaitId)
	request.TypeId = uint(u.CFG.Type.JoinTeamId)

	u.DB.Save(&request)
	c.JSON(http.StatusOK, request)
}

func (u *UserController) CancelRequest(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	authedUser, _ := c.Get("authedUser")
	user := authedUser.(middleware.AuthedUser)

	var userRequests []database.Request
	u.DB.Where("\"from\" = ?", user.Id).Find(&userRequests)

	if len(userRequests) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "You don't have any requests yet"})
		return
	}

	canceledRequest := false
	for _, n := range userRequests {
		if id == int(n.Id) {
			canceledRequest = true
			n.StatusId = uint(u.CFG.Status.CancelId)
			u.DB.Save(&n)
			c.JSON(http.StatusOK, n)
			return
		}
	}

	if canceledRequest == false {
		c.JSON(http.StatusNotFound, gin.H{"error": "You don't have request with this id"})
	}
	
	return
}

func (u *UserController) GetAllManagers(c *gin.Context) {
	var managers []database.User

	u.DB.Where("role = ?", u.CFG.Role.ManagerId).Find(&managers)

	if len(managers) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Manegers not found"})
		return
	}

	c.JSON(http.StatusOK, managers)
}

func (u *UserController) GetAllTeams(c *gin.Context) {
	var teams []database.Team

	u.DB.Find(&teams)

	if len(teams) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teams not found"})
		return
	}
	
	c.JSON(http.StatusOK, teams)
}