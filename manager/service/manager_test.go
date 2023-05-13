package service_test

import (
	"errors"
	"testing"

	"github.com/makovii/group_organiser/database"
	"github.com/makovii/group_organiser/manager/service"
	"github.com/stretchr/testify/assert"
)

type ManagerRepositoryMock struct {
	CreateTeamFunc  func(name string, managerID uint) (*database.Team, error)
	GetAllTeamsFunc func(managerID uint) (*[]database.Team, error)
	GetTeamFunc     func(teamID uint, managerID uint) (*database.Team, error)
	UpdateTeamFunc  func(teamID uint, managerID uint, name string) (*database.Team, error)
	DeleteTeamFunc  func(teamID uint, managerID uint) error
}

func (m *ManagerRepositoryMock) CreateTeam(name string, managerID uint) (*database.Team, error) {
	return m.CreateTeamFunc(name, managerID)
}

func (m *ManagerRepositoryMock) GetAllTeams(managerID uint) (*[]database.Team, error) {
	return m.GetAllTeamsFunc(managerID)
}

func (m *ManagerRepositoryMock) GetTeam(teamID uint, managerID uint) (*database.Team, error) {
	return m.GetTeamFunc(teamID, managerID)
}

func (m *ManagerRepositoryMock) UpdateTeam(teamID uint, managerID uint, name string) (*database.Team, error) {
	return m.UpdateTeamFunc(teamID, managerID, name)
}

func (m *ManagerRepositoryMock) DeleteTeam(teamID uint, managerID uint) error {
	return m.DeleteTeamFunc(teamID, managerID)
}

func TestCreateTeam(t *testing.T) {
	// given
	managerID := uint(1)
	name := "team name"

	expectedTeam := &database.Team{
		Id:        1,
		Name:      name,
		ManagerID: managerID,
	}

	mockRepo := &ManagerRepositoryMock{
		CreateTeamFunc: func(name string, managerID uint) (*database.Team, error) {
			return expectedTeam, nil
		},
	}

	service := service.NewManagerService(mockRepo)

	// when
	team, err := service.CreateTeam(name, managerID)

	// then
	assert.NoError(t, err)
	assert.Equal(t, expectedTeam, team)
}

func TestCreateTeamError(t *testing.T) {
	// given
	managerID := uint(1)
	name := "team name"

	mockRepo := &ManagerRepositoryMock{
		CreateTeamFunc: func(name string, managerID uint) (*database.Team, error) {
			return nil, errors.New("error creating team")
		},
	}

	service := service.NewManagerService(mockRepo)

	// when
	team, err := service.CreateTeam(name, managerID)

	// then
	assert.Error(t, err)
	assert.Nil(t, team)
}
