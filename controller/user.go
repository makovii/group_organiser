package controller

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/makovii/group_organiser/config"
	"github.com/makovii/group_organiser/database"
	"gorm.io/gorm"
)

type IUserService interface {
	GetUserById(id int) (*database.User, error)
	GetNotifications(to int) (*[]database.Request, error)
	JoinTeam(c*gin.Context, CFG *config.Config, body BodyJoinTeam) (*database.Request, error)
	LeaveTeam(c*gin.Context, CFG *config.Config, body BodyJoinTeam) (*database.Request, error)
	CancelRequest(c*gin.Context, CFG *config.Config, id int) (*database.Request, error)
	GetAllManagers() (*[]database.User, error)
	GetAllTeams() (*[]database.Team, error)
}

type UserController struct {
	DB *gorm.DB
	CFG *config.Config
	service IUserService
}

type BodyJoinTeam struct {
	TeamId	uint	`json:"teamId"`
}

func NewUserController(db *gorm.DB, cfg *config.Config, service IUserService) *UserController {
	return &UserController{ DB: db, CFG: cfg, service: service }
}

func (u *UserController) GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	
	var user *database.User

	user, err := u.service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *UserController) GetNotifications(c *gin.Context) {
	to, _ := strconv.Atoi(c.Query("to"))

	var requests *[]database.Request

	requests, err := u.service.GetNotifications(to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, requests)
}

func (u *UserController) JoinTeam(c *gin.Context) {
	var body BodyJoinTeam
	
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request, err := u.service.JoinTeam(c, u.CFG, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, request)
}

func (u *UserController) LeaveTeam(c *gin.Context) {
	var body BodyJoinTeam
	
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request, err := u.service.LeaveTeam(c, u.CFG, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, request)
}

func (u *UserController) CancelRequest(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	request, err := u.service.CancelRequest(c, u.CFG, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, request)		
	}
}

func (u *UserController) GetAllManagers(c *gin.Context) {

	managers, err := u.service.GetAllManagers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} 

	if len(*managers) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Manegers not found"})
		return
	}

	c.JSON(http.StatusOK, managers)
}

func (u *UserController) GetAllTeams(c *gin.Context) {
	teams, err := u.service.GetAllTeams()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} 

	if len(*teams) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teams not found"})
		return
	}

	c.JSON(http.StatusOK, teams)
}
