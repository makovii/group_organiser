package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/makovii/group_organiser/database"
	"github.com/makovii/group_organiser/middleware"
	"gorm.io/gorm"
)

type ManagerController struct {
	DB *gorm.DB
}

func NewManagerController(db *gorm.DB) *ManagerController {
	return &ManagerController{DB: db}
}

func (m *ManagerController) CreateTeam(c *gin.Context) {
	var team database.Team
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authedUser, _ := c.Get("authedUser")
	user := authedUser.(middleware.AuthedUser)
	
	team.ManagerID = uint(user.Id)

	if err := m.DB.Create(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, team)
}

func (m *ManagerController) GetAllTeams(c *gin.Context) {
	var teams []database.Team
	if err := m.DB.Find(&teams).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teams)
}

func (m *ManagerController) GetTeam(c *gin.Context) {
	id := c.Param("id")

	var team database.Team
	if err := m.DB.Where("id = ?", id).First(&team).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "team not found"})
		return
	}

	c.JSON(http.StatusOK, team)
}

func (m *ManagerController) UpdateTeam(c *gin.Context) {
	id := c.Param("id")

	var team database.Team
	if err := m.DB.Where("id = ?", id).First(&team).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "team not found"})
		return
	}

	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := m.DB.Model(&team).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, team)
}

func (m *ManagerController) DeleteTeam(c *gin.Context) {
	id := c.Param("id")

	var team database.Team
	if err := m.DB.Where("id = ?", id).First(&team).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "team not found"})
		return
	}

	if err := m.DB.Delete(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "team deleted"})
}

func (m *ManagerController) AcceptUserRequest(c *gin.Context) {
	teamID := c.Param("team_id")
	userIDStr := c.Param("user_id")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	var team database.Team
	if err := m.DB.Where("id = ?", teamID).First(&team).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "team not found"})
		return
	}

	var user database.User
	if err := m.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if !team.HasUserRequest(uint(userID)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user has no request to join this team"})
		return
	}

	if err := team.AcceptUserRequest(uint(userID), m.DB); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "user request accepted"})
	}
}
