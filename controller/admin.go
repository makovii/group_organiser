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

func (u UserController) GetAdmin(c *gin.Context) {
	admin := "admin1"
	c.JSON(http.StatusOK, admin)
	return
}

func (u UserController) GetById(c *gin.Context) {
	GetById := "getId"
	c.JSON(http.StatusOK, GetById)
	return
}

func (u UserController) BanById(c *gin.Context) {
	BanById := "banId"
	c.JSON(http.StatusOK, BanById)
	return
}

func (u UserController) GetTeans(c *gin.Context) {
	GetTeams := "teams1"
	c.JSON(http.StatusOK, GetTeams)
	return
}
