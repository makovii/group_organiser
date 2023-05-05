package repository

import (
	"github.com/makovii/group_organiser/database"
	"gorm.io/gorm"
)

type ManagerRepository struct {
	DB *gorm.DB
}

func NewManagerRepository(db *gorm.DB) *ManagerRepository {
	return &ManagerRepository{DB: db}
}

func (mr *ManagerRepository) CreateTeam(name string, managerID uint) (*database.Team, error) {
	var team database.Team
	team.ManagerID = managerID
	result := mr.DB.Create(&team)
	if result.Error != nil {
		return nil, result.Error
	}
	return &team, nil
}

func (mr *ManagerRepository) GetAllTeams(managerID uint) (*[]database.Team, error) {
	var teams []database.Team
	result := mr.DB.Where("manager_id = ?", managerID).Find(&teams)
	if result.Error != nil {
		return nil, result.Error
	}
	return &teams, nil
}

func (mr *ManagerRepository) GetTeam(teamID uint, managerID uint) (*database.Team, error) {
	var team database.Team
	if err := mr.DB.Where("id = ? AND manager_id = ?", teamID, managerID).First(&team).Error; err != nil {
		return &team, err
	}
	return &team, nil
}

func (mr *ManagerRepository) UpdateTeam(teamID uint, managerID uint, name string) (*database.Team, error) {
	var team database.Team
	if err := mr.DB.Where("id = ?", teamID).First(&team).Error; err != nil {
		return nil, err
	}
	team.Name = name
	team.ManagerID = managerID
	result := mr.DB.Save(team)
	if result.Error != nil {
		return nil, result.Error
	}
	return &team, nil
}

func (mr *ManagerRepository) DeleteTeam(teamID uint, managerID uint) error {
	result := mr.DB.Where("id = ? AND manager_id = ?", teamID, managerID).Delete(&database.Team{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
