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

type MockAdminController struct{}

func (s *MockAdminController) GetAdmins() (*[]database.User, error) {
	user := database.User{
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
	users := &[]database.User{user}

	return users, nil
}

func (s *MockAdminController) GetUserById(id int) (*database.User, error) {
	user := &database.User{
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

	return user, nil
}

func (s *MockAdminController) GetTeamById(id uint) (*database.Team, error) {
	managerID := uint(1)
	name := "team name"

	team := &database.Team{
		Id:        1,
		Name:      name,
		ManagerID: managerID,
	}

	return team, nil
}

func (s *MockAdminController) GetTeams() (*[]database.Team, error) {
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

func (s *MockAdminController) BanById(id int) (*database.User, error) {
	user := &database.User{
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
	return user, nil
}

func (s *MockAdminController) AcceptManagerRegistration(id int) (*database.User, error) {
	user := &database.User{
		Id:           3,
		Name:         "name",
		Email:        "email@gmail.com",
		Password:     "password",
		Request:      []database.Request{},
		Notification: []string{},
		Teams:        []database.Team{},
		Ban:          false,
		Role:         2,
	}
	return user, nil
}

func setupAdminRouter(service IAdminService) *gin.Engine {
	router := gin.Default()
	controller := NewAdminController(service)

	router.GET("/getAdmins", controller.GetAdmins)
	router.GET("/getUserById", controller.GetUserById)
	router.GET("/getTeamById", controller.GetTeamById)
	router.POST("/banById", controller.BanById)
	router.POST("/getTeams", controller.GetTeams)
	router.POST("/acceptRegistration", controller.AcceptManagerRegistration)

	return router
}

func TestGetAdmins(t *testing.T) {
	service := &MockAdminController{}
	router := setupAdminRouter(service)

	w := httptest.NewRecorder()
	body := BodyJoinTeam{}
	bodyBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest("GET", "/getAdmins", bytes.NewReader(bodyBytes))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	user := database.User{
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
	expectedUsers := []database.User{user}

	var responseAdmins []database.User
	err := json.Unmarshal(w.Body.Bytes(), &responseAdmins)
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, responseAdmins)
}

func TestGetUserById(t *testing.T) {
	service := &MockAdminController{}
	router := setupAdminRouter(service)

	w := httptest.NewRecorder()
	body := BodyCreateTeam{}
	bodyBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest("GET", "/getUserById?id=1", bytes.NewReader(bodyBytes))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseUser database.User
	err := json.Unmarshal(w.Body.Bytes(), &responseUser)
	assert.NoError(t, err)
	assert.Equal(t, "name", responseUser.Name)
}

func TestGetTeamById(t *testing.T) {
	service := &MockAdminController{}
	router := setupAdminRouter(service)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/getTeamById?id=1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var responseTeam database.Team
	err := json.Unmarshal(w.Body.Bytes(), &responseTeam)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), responseTeam.Id)
	assert.Equal(t, "team name", responseTeam.Name)
}

func TestGetTeams(t *testing.T) {
	service := &MockAdminController{}
	router := setupAdminRouter(service)

	w := httptest.NewRecorder()
	body := BodyJoinTeam{}
	bodyBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/getTeams", bytes.NewReader(bodyBytes))
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

func TestBanById(t *testing.T) {
	service := &MockAdminController{}
	router := setupAdminRouter(service)

	w := httptest.NewRecorder()
	body := BodyJoinTeam{}
	bodyBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/banById?id=1", bytes.NewReader(bodyBytes))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseUser database.User
	err := json.Unmarshal(w.Body.Bytes(), &responseUser)
	assert.NoError(t, err)
	assert.Equal(t, true, responseUser.Ban)

}

func TestAcceptManagerRegistration(t *testing.T) {
	service := &MockAdminController{}
	router := setupAdminRouter(service)

	w := httptest.NewRecorder()
	body := BodyJoinTeam{}
	bodyBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/acceptRegistration?id=3", bytes.NewReader(bodyBytes))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseUser database.User
	err := json.Unmarshal(w.Body.Bytes(), &responseUser)
	assert.NoError(t, err)
	assert.Equal(t, false, responseUser.Ban)
}
