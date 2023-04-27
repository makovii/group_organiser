package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/makovii/group_organiser/database"
	"gorm.io/gorm"
)

type AdminController struct {
	DB *gorm.DB
}

func NewAdminController(db *gorm.DB) *AdminController {
	return &AdminController{DB: db}
}

func (a *AdminController) GetAdmin(c *gin.Context) {
	var admin database.User
	c.JSON(http.StatusOK, admin)

}

func (a *AdminController) GetById(c *gin.Context) {
	var GetById database.User
	c.JSON(http.StatusOK, GetById)

}

func (a *AdminController) BanById(c *gin.Context) {
	var BanById database.User
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

	manager.Ban = false;

	a.DB.Save(&manager)
}
