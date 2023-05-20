package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/makovii/group_organiser/database"
)

type IAdminService interface {
	GetAdmins() (*[]database.User, error)
	GetUserById(id int) (*database.User, error)
	GetTeamById(id uint) (*database.Team, error)
	GetTeams() (*[]database.Team, error)
	BanById(id int) (*database.User, error)
	AcceptManagerRegistration(id int) (*database.User, error)
}

type AdminController struct {
	service IAdminService
}

func NewAdminController(service IAdminService) *AdminController {
	return &AdminController{service: service}
}

func (a *AdminController) GetAdmins(c *gin.Context) {
	admins, err := a.service.GetAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if len(*admins) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admins not found"})
		return
	}

	c.JSON(http.StatusOK, admins)

}

func (a *AdminController) GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var user *database.User

	user, err := a.service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, user)

}

func (a *AdminController) GetTeamById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var teams *database.Team

	teams, err := a.service.GetTeamById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, teams)
	}

}

func (a *AdminController) GetTeams(c *gin.Context) {
	teams, err := a.service.GetTeams()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if len(*teams) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teams not found"})
		return
	}

	c.JSON(http.StatusOK, teams)
}

func (a *AdminController) BanById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var BanById *database.User

	BanById, err := a.service.BanById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, BanById)
	}
}

func (a *AdminController) AcceptManagerRegistration(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var manager *database.User

	manager, err := a.service.AcceptManagerRegistration(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, manager)
	}
}
