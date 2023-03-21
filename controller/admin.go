package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminController struct {
	DB *gorm.DB
}

func NewAdminController(db *gorm.DB) *AdminController {
	return &AdminController{DB: db}
}

func (a AdminController) GetAdmin(c *gin.Context) {
	admin := "admin1"
	c.JSON(http.StatusOK, admin)

}

func (u AdminController) GetById(c *gin.Context) {
	GetById := "getId"
	c.JSON(http.StatusOK, GetById)

}

func (u AdminController) BanById(c *gin.Context) {
	BanById := "banId"
	c.JSON(http.StatusOK, BanById)

}

func (u AdminController) GetTeams(c *gin.Context) {
	GetTeams := "teams1"
	c.JSON(http.StatusOK, GetTeams)
}
