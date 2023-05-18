package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/makovii/group_organiser/database"
	"github.com/stretchr/testify/assert"
)

type MocUserController struct{
	GetUserByIdFunc					func (id int) (*database.User, error)
	GetNotificationsFunc		func(to int) (*[]database.Request, error)
	GetTeamByIdFunc					func(id uint) (*database.Team)
	CreateRequestFunc				func(request database.Request) (uint, error)
	SaveRequestFunc					func(request database.Request) (uint, error)
	GetRequestsByFromFunc		func(from int) (*[]database.Request, error)
	GetAllManagersFunc			func() (*[]database.User, error)
	GetAllTeamsFunc					func() (*[]database.Team, error)
	JoinTeamFunc						func(c *gin.Context, body BodyJoinTeam) (*database.Request, error)
	LeaveTeamFunc						func(c *gin.Context, body BodyJoinTeam) (*database.Request, error)
	CancelRequestFunc				func(c *gin.Context,id int) (*database.Request, error)
}

func (s *MocUserController) CancelRequest(c *gin.Context,id int) (*database.Request, error) {
	team := &database.Request{
		Id: 1,
		From: 2,
		To: 3,
		StatusId: 1,
		TypeId: 3,
	}

	return team, nil
}

func (m *MocUserController) GetUserById(id int) (*database.User, error) {
	user := &database.User{
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

	return user, nil
}

func (m *MocUserController) JoinTeam(c *gin.Context, body BodyJoinTeam) (*database.Request, error) {
	request := &database.Request{
		Id: 1,
		From: 2,
		To: 1,
		StatusId: 1,
		TypeId: 2,
	}
	return request, nil
}

func (m *MocUserController) LeaveTeam(c *gin.Context, body BodyJoinTeam) (*database.Request, error) {
	request := &database.Request{
		Id: 1,
		From: 2,
		To: 1,
		StatusId: 1,
		TypeId: 3,
	}
	return request, nil
}

func (m *MocUserController) GetNotifications(to int) (*[]database.Request, error) {
	notification := database.Request{
		Id: 1,
		From: 2,
		To: 3,
		StatusId: 1,
		TypeId: 3,
	}
	notifications := &[]database.Request{notification}

	return notifications, nil
}

func (m *MocUserController) GetAllManagers() (*[]database.User, error) {
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
	users := &[]database.User{user}

	return users, nil
}

func (m *MocUserController) GetAllTeams() (*[]database.Team, error) {
	managerID := uint(1)
	name := "team name"

	team := database.Team{
		Id:        1,
		Name:      name,
		ManagerID: managerID,
	}

	teams := &[]database.Team{team}
	
	return teams, nil
}

func setupUserRouter(service IUserService) *gin.Engine {
	router := gin.Default()
	controller := NewUserController(service)

	router.GET("/getUser", controller.GetUserById)
	router.GET("/myNotifications", controller.GetNotifications)
	router.POST("/joinTeam", controller.JoinTeam)
	router.POST("/leaveTeam", controller.LeaveTeam)
	router.POST("/cancelRequest", controller.CancelRequest)
	router.GET("/getAllManagers", controller.GetAllManagers)
	router.GET("/getAllTeams", controller.GetAllTeams)

	return router
}

func TestUserGetUserById(t *testing.T) {
	service := &MocUserController{}
	router := setupUserRouter(service)

	w := httptest.NewRecorder()
	body := BodyCreateTeam{}
	bodyBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest("GET", "/getUser?id=1", bytes.NewReader(bodyBytes))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseUser database.User
	err := json.Unmarshal(w.Body.Bytes(), &responseUser)
	assert.NoError(t, err)
	assert.Equal(t, "name", responseUser.Name)
}

func TestUserMyNotifications(t *testing.T) {
	service := &MocUserController{}
	router := setupUserRouter(service)

	w := httptest.NewRecorder()
	body := BodyCreateTeam{}
	bodyBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest("GET", "/myNotifications?to=1", bytes.NewReader(bodyBytes))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	notification := database.Request{
		Id: 1,
		From: 0,
		To: 0,
		StatusId: 1,
		TypeId: 3,
	}
	expectedNotification := []database.Request{notification}

	var response []database.Request
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedNotification, response)
}

func TestUserJoinTeam(t *testing.T) {
	service := &MocUserController{}
	router := setupUserRouter(service)

	w := httptest.NewRecorder()
	body := BodyJoinTeam{
		TeamId: 1,
	}
	bodyBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/joinTeam", bytes.NewReader(bodyBytes))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	notification := database.Request{
		Id: 1,
		From: 0,
		To: 0,
		StatusId: 1,
		TypeId: 2,
	}

	var response database.Request
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, notification, response)
}

func TestUserLeaveTeam(t *testing.T) {
	service := &MocUserController{}
	router := setupUserRouter(service)

	w := httptest.NewRecorder()
	body := BodyJoinTeam{
		TeamId: 1,
	}
	bodyBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/leaveTeam", bytes.NewReader(bodyBytes))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	notification := database.Request{
		Id: 1,
		From: 0,
		To: 0,
		StatusId: 1,
		TypeId: 3,
	}

	var response database.Request
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, notification, response)
}

func TestUserCancelRequest(t *testing.T) {
	service := &MocUserController{}
	router := setupUserRouter(service)

	w := httptest.NewRecorder()
	body := BodyJoinTeam{}
	bodyBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/cancelRequest?id=1", bytes.NewReader(bodyBytes))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	notification := database.Request{
		Id: 1,
		From: 0,
		To: 0,
		StatusId: 1,
		TypeId: 3,
	}

	var response database.Request
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, notification, response)
}

func TestUserGetAllManagers(t *testing.T) {
	service := &MocUserController{}
	router := setupUserRouter(service)

	w := httptest.NewRecorder()
	body := BodyJoinTeam{}
	bodyBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest("GET", "/getAllManagers", bytes.NewReader(bodyBytes))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

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
	expectedUsers := []database.User{user}

	var response []database.User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, response)
}

func TestUserGetAllTeams(t *testing.T) {
	service := &MocUserController{}
	router := setupUserRouter(service)

	w := httptest.NewRecorder()
	body := BodyJoinTeam{}
	bodyBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest("GET", "/getAllTeams", bytes.NewReader(bodyBytes))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	managerID := uint(1)
	name := "team name"

	team := database.Team{
		Id:        1,
		Name:      name,
		ManagerID: managerID,
	}

	expectedTeams := []database.Team{team}

	var response []database.Team
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedTeams, response)
}
