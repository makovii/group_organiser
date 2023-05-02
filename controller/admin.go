package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/makovii/group_organiser/config"
	"github.com/makovii/group_organiser/database"
	"gorm.io/gorm"
)

type AdminController struct {
	DB  *gorm.DB
	CFG *config.Config
}

func NewAdminController(db *gorm.DB, cfg *config.Config) *AdminController {
	return &AdminController{DB: db, CFG: cfg}
}

func (a *AdminController) GetAdmins(c *gin.Context) {
	var admins []database.User

	a.DB.Where("role = ?", a.CFG.Role.AdminId).Find(&admins)

	if len(admins) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admins"})
		return
	}

	c.JSON(http.StatusOK, admins)

}

func (a *AdminController) GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var user database.User

	if a.DB.Where("id = ?", id).First(&user) != nil {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
	}
}

func (a *AdminController) GetTeamById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var teams database.Team

	if a.DB.Where("id = ?", id).First(&teams) != nil {
		c.JSON(http.StatusOK, teams)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
	}

}

func (a *AdminController) BanById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	var BanById database.User
	a.DB.Where("id = ?", id).First(&BanById)
	BanById.Ban = true
	c.JSON(http.StatusOK, BanById)

}

func (a *AdminController) GetTeams(c *gin.Context) {
	var teams []database.Team
	if err := a.DB.Find(&teams).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teams)
}

func (a *AdminController) AcceptManagerRegistration(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))

	var manager database.User
	a.DB.Where("id = ?", id).First(&manager)

	manager.Ban = false

	a.DB.Save(&manager)
	c.JSON(http.StatusOK, manager)
}
