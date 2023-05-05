package service

import (
	"errors"

	"github.com/makovii/group_organiser/database"
)

type IManagerRepository interface {
	CreateTeam(name string, managerID uint) (*database.Team, error)
	GetAllTeams(managerID uint) (*[]database.Team, error)
	GetTeam(teamID uint, managerID uint) (*database.Team, error)
	UpdateTeam(teamID uint, managerID uint, name string) (*database.Team, error)
	DeleteTeam(teamID uint, managerID uint) error
}

type ManagerService struct {
	managerRepository IManagerRepository
}

func NewManagerService(managerRepository IManagerRepository) *ManagerService {
	return &ManagerService{managerRepository}
}

func (ms *ManagerService) CreateTeam(name string, managerID uint) (*database.Team, error) {

	team, err := ms.managerRepository.CreateTeam(name, managerID)
	if err != nil {
		return nil, errors.New("Team not found")
	}
	return team, nil
}

func (ms *ManagerService) GetAllTeams(managerID uint) (*[]database.Team, error) {
	team, err := ms.managerRepository.GetAllTeams(managerID)
	if err != nil {
		return nil, errors.New("Failed to get teams")
	}
	return team, nil
}

func (ms *ManagerService) GetTeam(teamID uint, managerID uint) (*database.Team, error) {
	team, err := ms.managerRepository.GetTeam(teamID, managerID)
	if err != nil {
		return nil, errors.New("Failed to get team")
	}
	return team, nil
}

func (ms *ManagerService) UpdateTeam(teamID uint, managerID uint, name string) (*database.Team, error) {
	team, err := ms.managerRepository.UpdateTeam(teamID, managerID, name)
	if err != nil {
		return nil, errors.New("Failed to update team")
	}
	return team, nil
}

func (ms *ManagerService) DeleteTeam(teamID uint, managerID uint) error {
	err := ms.managerRepository.DeleteTeam(teamID, managerID)
	if err != nil {
		return errors.New("Failed to delete team")
	}
	return nil
}
