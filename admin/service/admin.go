package service

import (
	"errors"

	"github.com/makovii/group_organiser/database"
)

type IAdminRepository interface {
	GetAdmins() (*[]database.User, error)
	GetUserById(id int) (*database.User, error)
	GetTeamById(id uint) *database.Team
	GetTeams() (*[]database.Team, error)
	BanById(id int) (*database.User, error)
	AcceptManagerRegistration(id int) (*database.User, error)
}

type AdminService struct {
	adminRepository IAdminRepository
}

func NewAdminService(adminRepository IAdminRepository) *AdminService {
	return &AdminService{adminRepository}
}

func (a *AdminService) GetAdmins() (*[]database.User, error) {
	admins, err := a.adminRepository.GetAdmins()

	if err != nil {
		return nil, errors.New("get Admins error")
	}

	return admins, nil
}

func (a *AdminService) GetUserById(id int) (*database.User, error) {
	return a.adminRepository.GetUserById(id)
}

func (a *AdminService) GetTeamById(id uint) (*database.Team, error) {
	team := a.adminRepository.GetTeamById(id)

	if team.Id == 0 {
		return nil, errors.New("team not found")
	}
	return team, nil
}

func (a *AdminService) GetTeams() (*[]database.Team, error) {
	teams, err := a.adminRepository.GetTeams()

	if err != nil {
		return nil, errors.New("get Teams error")
	}
	return teams, nil
}

func (a *AdminService) BanById(id int) (*database.User, error) {
	return a.adminRepository.BanById(id)
}

func (a *AdminService) AcceptManagerRegistration(id int) (*database.User, error) {
	return a.adminRepository.AcceptManagerRegistration(id)
}
