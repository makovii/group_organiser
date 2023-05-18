package service_test

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/makovii/group_organiser/config"
	"github.com/makovii/group_organiser/controller"
	"github.com/makovii/group_organiser/database"
	"github.com/makovii/group_organiser/middleware"
	"github.com/makovii/group_organiser/user/service"
	"github.com/stretchr/testify/assert"
)

type UserServiceMock struct {
	GetUserByIdFunc					func (id int) (*database.User, error)
	GetNotificationsFunc		func(to int) (*[]database.Request, error)
	GetTeamByIdFunc					func(id uint) (*database.Team)
	CreateRequestFunc				func(request database.Request) (uint, error)
	SaveRequestFunc					func(request database.Request) (uint, error)
	GetRequestsByFromFunc		func(from int) (*[]database.Request, error)
	GetAllManagersFunc			func() (*[]database.User, error)
	GetAllTeamsFunc					func() (*[]database.Team, error)
	JoinTeamFunc						func(c *gin.Context, body controller.BodyJoinTeam) (*database.Request, error)
	LeaveTeamFunc						func(c *gin.Context, body controller.BodyJoinTeam) (*database.Request, error)
	CancelRequestFunc				func(c *gin.Context,id int) (*database.Request, error)
}

func (m *UserServiceMock) GetUserById(id int) (*database.User, error) {
	return m.GetUserByIdFunc(id)
}

func (m *UserServiceMock) GetNotifications(to int) (*[]database.Request, error) {
	return m.GetNotificationsFunc(to)
}

func (m *UserServiceMock) GetTeamById(id uint) (*database.Team) {
	return m.GetTeamByIdFunc(id)
}

func (m *UserServiceMock) CreateRequest(request database.Request) (uint, error) {
	return m.CreateRequestFunc(request)
}
func (m *UserServiceMock) SaveRequest(request database.Request) (uint, error) {
	return m.SaveRequestFunc(request)
}

func (m *UserServiceMock) GetRequestsByFrom(from int) (*[]database.Request, error) {
	return m.GetRequestsByFromFunc(from)
}

func (m *UserServiceMock) GetAllManagers() (*[]database.User, error) {
	return m.GetAllManagersFunc()
}

func (m *UserServiceMock) GetAllTeams() (*[]database.Team, error) {
	return m.GetAllTeamsFunc()
}

func TestGetUserById(t *testing.T) {
	// given

	expectedUser := &database.User{
		Id:       		1,
		Name:     		"name",
		Email:   			"email@gmail.com",
		Password: 		"password",
		Request:  		[]database.Request{},
		Notification: []string{},
		Teams:				[]database.Team{},
		Ban: 					false,
		Role:					3,
	}

	mockRepo := &UserServiceMock{
		GetUserByIdFunc: func(id int) (*database.User, error) {
			return expectedUser, nil
		},
	}
	cfg := config.GetConfig()
	service := service.NewUserService(cfg, mockRepo)

	// when
	user, err := service.GetUserById(1)

	// then
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestGetNotifications(t *testing.T) {
	// given
	notification := database.Request{
		Id: 1,
		From: 2,
		To: 3,
		StatusId: 1,
		TypeId: 3,
	}
	expectedNotification := []database.Request{notification}

	mockRepo := &UserServiceMock{
		GetNotificationsFunc: func(to int) (*[]database.Request, error) {
			return &expectedNotification, nil
		},
	}

	cfg := config.GetConfig()
	service := service.NewUserService(cfg, mockRepo)

	// when
	result, err := service.GetNotifications(2)

	// then
	assert.NoError(t, err)
	assert.Equal(t, &expectedNotification, result)
}

func TestJoinTeam(t *testing.T) {
	// given
	managerID := uint(1)
	name := "team name"

	expectedTeam := &database.Team{
		Id:        1,
		Name:      name,
		ManagerID: managerID,
	}

	expectedRequest := &database.Request{
		Id: 1,
		From: 2,
		To: 1,
		StatusId: 1,
		TypeId: 2,
	}

	mockRepo := &UserServiceMock{
		JoinTeamFunc: func(c *gin.Context, body controller.BodyJoinTeam) (*database.Request, error) {
			return expectedRequest, nil
		},
		GetTeamByIdFunc: func (id uint) (*database.Team) {
			return expectedTeam
		},
		CreateRequestFunc: func(request database.Request) (uint, error) {
			return 1, nil
		},
	}

	cfg := config.GetConfig()
	service := service.NewUserService(cfg, mockRepo)

	// when
	gin.SetMode(gin.TestMode)

	authedUser := middleware.AuthedUser{
		Id:             2,
		Name:           "John",
		Email:          "john@doe.com",
		Role: 					2,
	}

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Set("authedUser", authedUser)
	body := controller.BodyJoinTeam{ TeamId: 3 }
	request, err := service.JoinTeam(c, body)

	// then
	assert.NoError(t, err)
	assert.Equal(t, expectedRequest, request)
}

func TestLeaveTeam(t *testing.T) {
	// given
	managerID := uint(1)
	name := "team name"

	expectedTeam := &database.Team{
		Id:        1,
		Name:      name,
		ManagerID: managerID,
	}

	expectedRequest := &database.Request{
		Id: 1,
		From: 2,
		To: 1,
		StatusId: 1,
		TypeId: 3,
	}

	mockRepo := &UserServiceMock{
		LeaveTeamFunc: func(c *gin.Context, body controller.BodyJoinTeam) (*database.Request, error) {
			return expectedRequest, nil
		},
		GetTeamByIdFunc: func (id uint) (*database.Team) {
			return expectedTeam
		},
		CreateRequestFunc: func(request database.Request) (uint, error) {
			return 1, nil
		},
	}

	cfg := config.GetConfig()
	service := service.NewUserService(cfg, mockRepo)

	// when
	gin.SetMode(gin.TestMode)

	authedUser := middleware.AuthedUser{
		Id:             2,
		Name:           "John",
		Email:          "john@doe.com",
		Role: 					2,
	}

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Set("authedUser", authedUser)
	body := controller.BodyJoinTeam{ TeamId: 3 }
	request, err := service.LeaveTeam(c, body)

	// then
	assert.NoError(t, err)
	assert.Equal(t, expectedRequest, request)
}

func TestGetAllManagers(t *testing.T) {
	// given
	user := database.User{
		Id:       		1,
		Name:     		"name",
		Email:   			"email@gmail.com",
		Password: 		"password",
		Request:  		[]database.Request{},
		Notification: []string{},
		Teams:				[]database.Team{},
		Ban: 					false,
		Role:					3,
	}
	expectedUsers := &[]database.User{user}

	mockRepo := &UserServiceMock{
		GetAllManagersFunc: func() (*[]database.User, error) {
			return expectedUsers, nil
		},
	}

	cfg := config.GetConfig()
	service := service.NewUserService(cfg, mockRepo)

	// when
	managers, err := service.GetAllManagers()

	// then
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, managers)
}

func TestCancelRequest(t *testing.T) {
	// given
	expectedRequest := &database.Request{
		Id: 1,
		From: 2,
		To: 3,
		StatusId: 1,
		TypeId: 3,
	}

	expectedRequest2 := database.Request{
		Id: 1,
		From: 2,
		To: 3,
		StatusId: 1,
		TypeId: 3,
	}

	expectedRequests := []database.Request{expectedRequest2}

	mockRepo := &UserServiceMock{
		CancelRequestFunc: func(c *gin.Context, id int) (*database.Request, error) {
			return expectedRequest, nil
		},
		GetRequestsByFromFunc:	func(from int) (*[]database.Request, error){
			return &expectedRequests, nil
		},
	}

	cfg := config.GetConfig()
	service := service.NewUserService(cfg, mockRepo)

	// when
		gin.SetMode(gin.TestMode)

	authedUser := middleware.AuthedUser{
		Id:             2,
		Name:           "John",
		Email:          "john@doe.com",
		Role: 					2,
	}
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Set("authedUser", authedUser)
	_, err := service.CancelRequest(c, 2)

	// then
	assert.Error(t, errors.New("You don't have request with this id"), err)
}

func TestGetAllTeams(t *testing.T) {
	// given
	managerID := uint(1)
	name := "team name"

	team := database.Team{
		Id:        1,
		Name:      name,
		ManagerID: managerID,
	}

	expectedTeams := []database.Team{team}

	mockRepo := &UserServiceMock{
		GetAllTeamsFunc: func() (*[]database.Team, error) {
			return &expectedTeams, nil
		},
	}

	cfg := config.GetConfig()
	service := service.NewUserService(cfg, mockRepo)

	// when
	teams, err := service.GetAllTeams()

	// then
	assert.NoError(t, err)
	assert.Equal(t, &expectedTeams, teams)
}

