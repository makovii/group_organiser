package repository

import (
	"errors"

	"github.com/makovii/group_organiser/config"
	"github.com/makovii/group_organiser/database"
	"gorm.io/gorm"
)

type AdminRepository struct {
	DB  *gorm.DB
	CFG *config.Config
}

func NewAdminRepository(db *gorm.DB, cfg *config.Config) *AdminRepository {
	return &AdminRepository{DB: db, CFG: cfg}
}

func (a *AdminRepository) GetAdmins() (*[]database.User, error) {
	var admins []database.User

	if err := a.DB.Where("role = ?", a.CFG.Role.AdminId).Find(&admins).Error; err != nil {
		return &admins, err
	}
	return &admins, nil
}

func (a *AdminRepository) GetUserById(id int) (*database.User, error) {
	var user database.User

	if err := a.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return &user, err
	}
	return &user, nil
}

func (a *AdminRepository) GetTeamById(id uint) *database.Team {
	var team database.Team

	a.DB.Where("id = ?", id).First(&team)
	return &team
}

func (a *AdminRepository) GetTeams() (*[]database.Team, error) {
	var teams []database.Team

	if err := a.DB.Find(&teams).Error; err != nil {
		return &teams, err
	}
	return &teams, nil
}

func (a *AdminRepository) BanById(id int) (*database.User, error) {
	var BanById database.User

	if err := a.DB.Model(&database.User{}).Where("id = ?", id).Update("ban", true).Error; err != nil {
		return nil, errors.New("manager update error")
	}
	if err := a.DB.Where("id = ?", id).First(&BanById).Error; err != nil {
		return nil, errors.New("manager find error")
	}
	return &BanById, nil

}

func (a *AdminRepository) AcceptManagerRegistration(id int) (*database.User, error) {
	var manager database.User

	if err := a.DB.Model(&database.User{}).Where("id = ?", id).Update("ban", false).Error; err != nil {
		return nil, errors.New("manager update error")
	}
	if err := a.DB.Where("id = ?", id).First(&manager).Error; err != nil {
		return nil, errors.New("manager find error")
	}
	return &manager, nil

}
