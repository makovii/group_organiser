package service_test

import (
	"testing"

	"github.com/makovii/group_organiser/admin/service"
	"github.com/makovii/group_organiser/database"
	"github.com/stretchr/testify/assert"
)

type AdminRepositoryMock struct {
	GetAdminsFunc                 func() (*[]database.User, error)
	GetUserByIdFunc               func(id int) (*database.User, error)
	GetTeamByIdFunc               func(id uint) *database.Team
	GetTeamsFunc                  func() (*[]database.Team, error)
	BanByIdFunc                   func(id int) (*database.User, error)
	AcceptManagerRegistrationFunc func(id int) (*database.User, error)
}

func (m *AdminRepositoryMock) GetAdmins() (*[]database.User, error) {
	return m.GetAdminsFunc()
}

func (m *AdminRepositoryMock) GetUserById(id int) (*database.User, error) {
	return m.GetUserByIdFunc(id)
}

func (m *AdminRepositoryMock) GetTeamById(id uint) *database.Team {
	return m.GetTeamByIdFunc(id)
}

func (m *AdminRepositoryMock) GetTeams() (*[]database.Team, error) {
	return m.GetTeamsFunc()
}

func (m *AdminRepositoryMock) BanById(id int) (*database.User, error) {
	return m.BanByIdFunc(id)
}

func (m *AdminRepositoryMock) AcceptManagerRegistration(id int) (*database.User, error) {
	return m.AcceptManagerRegistrationFunc(id)
}

func TestGetAdmins(t *testing.T) {
	// given
	user := database.User{
		Id:           1,
		Name:         "name",
		Email:        "email@gmail.com",
		Password:     "password",
		Request:      []database.Request{},
		Notification: []string{},
		Teams:        []database.Team{},
		Ban:          false,
		Role:         1,
	}
	expectedUsers := &[]database.User{user}

	mockRepo := &AdminRepositoryMock{
		GetAdminsFunc: func() (*[]database.User, error) {
			return expectedUsers, nil
		},
	}

	service := service.NewAdminService(mockRepo)

	// when
	managers, err := service.GetAdmins()

	// then
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, managers)
}

func TestGetUserById(t *testing.T) {
	// given

	expectedUser := &database.User{
		Id:           1,
		Name:         "name",
		Email:        "email@gmail.com",
		Password:     "password",
		Request:      []database.Request{},
		Notification: []string{},
		Teams:        []database.Team{},
		Ban:          false,
		Role:         3,
	}

	mockRepo := &AdminRepositoryMock{
		GetUserByIdFunc: func(id int) (*database.User, error) {
			return expectedUser, nil
		},
	}

	service := service.NewAdminService(mockRepo)

	// when
	user, err := service.GetUserById(1)

	// then
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestGetTeamById(t *testing.T) {
	// given
	managerID := uint(1)
	name := "name"

	expectedTeam := &database.Team{
		Id:        1,
		Name:      name,
		ManagerID: managerID,
	}

	mockRepo := &AdminRepositoryMock{
		GetTeamByIdFunc: func(id uint) *database.Team {
			return expectedTeam
		},
	}

	service := service.NewAdminService(mockRepo)
	// when
	team, err := service.GetTeamById(1)
	//then
	assert.NoError(t, err)
	assert.Equal(t, expectedTeam, team)
}

func TestGetTeams(t *testing.T) {
	// given
	managerID := uint(1)
	name := "name"

	team := database.Team{
		Id:        1,
		Name:      name,
		ManagerID: managerID,
	}

	expectedTeams := []database.Team{team}

	mockRepo := &AdminRepositoryMock{
		GetTeamsFunc: func() (*[]database.Team, error) {
			return &expectedTeams, nil
		},
	}

	service := service.NewAdminService(mockRepo)

	// when
	teams, err := service.GetTeams()

	// then
	assert.NoError(t, err)
	assert.Equal(t, &expectedTeams, teams)
}

func TestBanById(t *testing.T) {
	// given

	expectedUser := &database.User{
		Id:           1,
		Name:         "name",
		Email:        "email@gmail.com",
		Password:     "password",
		Request:      []database.Request{},
		Notification: []string{},
		Teams:        []database.Team{},
		Ban:          true,
		Role:         3,
	}

	mockRepo := &AdminRepositoryMock{
		BanByIdFunc: func(id int) (*database.User, error) {
			return expectedUser, nil
		},
	}

	service := service.NewAdminService(mockRepo)

	// when
	user, err := service.BanById(1)

	// then
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestAcceptManagerRegistration(t *testing.T) {
	// given

	expectedUser := &database.User{
		Id:           1,
		Name:         "name",
		Email:        "email@gmail.com",
		Password:     "password",
		Request:      []database.Request{},
		Notification: []string{},
		Teams:        []database.Team{},
		Ban:          false,
		Role:         2,
	}

	mockRepo := &AdminRepositoryMock{
		AcceptManagerRegistrationFunc: func(id int) (*database.User, error) {
			return expectedUser, nil
		},
	}

	service := service.NewAdminService(mockRepo)

	// when
	user, err := service.AcceptManagerRegistration(1)

	// then
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}
