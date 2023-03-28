package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/makovii/group_organiser/database"
)

type AdminController struct {
	DB *gorm.DB
}

func NewAdminController(db *gorm.DB) *AdminController {
	return &AdminController{DB: db}
}

func (a *AdminController) GetAdmin(c *gin.Context) {
	var admin database.Admin
	c.JSON(http.StatusOK, admin)

}

func (a *AdminController) GetById(c *gin.Context) {
	var GetById database.Admin
	c.JSON(http.StatusOK, GetById)

}

func (a *AdminController) BanById(c *gin.Context) {
	var BanById database.Player
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
