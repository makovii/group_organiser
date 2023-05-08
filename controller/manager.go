package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/makovii/group_organiser/config"
	"github.com/makovii/group_organiser/database"
	"gorm.io/gorm"
)

type IManagerService interface {
	CreateTeam(name string, managerID uint) (*database.Team, error)
	GetAllTeams(managerID uint) (*[]database.Team, error)
	GetTeam(teamID uint, managerID uint) (*database.Team, error)
	UpdateTeam(teamID uint, managerID uint, name string) (*database.Team, error)
	DeleteTeam(teamID uint, managerID uint) error
}

type ManagerController struct {
	Service IManagerService
}

type BodyCreateTeam struct {
	Name string `json:"name"`
}

type BodyUpdateTeam struct {
	Name string `json:"name"`
}

func NewManagerController(db *gorm.DB, cfg *config.Config, service IManagerService) *ManagerController {
	return &ManagerController{Service: service}
}

func (mc *ManagerController) CreateTeam(c *gin.Context) {
	var body BodyCreateTeam

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	managerID, err := strconv.Atoi(c.Query("managerId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team, err := mc.Service.CreateTeam(body.Name, uint(managerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, team)
}

func (mc *ManagerController) GetAllTeams(c *gin.Context) {
	managerID, err := strconv.Atoi(c.Query("managerId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teams, err := mc.Service.GetAllTeams(uint(managerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(*teams) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teams not found"})
		return
	}

	c.JSON(http.StatusOK, teams)
}

func (mc *ManagerController) GetTeam(c *gin.Context) {
	managerID, err := strconv.Atoi(c.Query("managerId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teamID, err := strconv.Atoi(c.Param("teamId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team, err := mc.Service.GetTeam(uint(teamID), uint(managerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, team)
}
func (mc *ManagerController) UpdateTeam(c *gin.Context) {
	var body BodyUpdateTeam

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	managerID, err := strconv.Atoi(c.Query("managerId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	teamID, err := strconv.Atoi(c.Param("teamId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team, err := mc.Service.UpdateTeam(uint(teamID), uint(managerID), body.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, team)
}
func (mc *ManagerController) DeleteTeam(c *gin.Context) {
	teamID, err := strconv.Atoi(c.Param("teamId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	managerID, err := strconv.Atoi(c.Query("managerId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = mc.Service.DeleteTeam(uint(teamID), uint(managerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Team deleted"})
}
