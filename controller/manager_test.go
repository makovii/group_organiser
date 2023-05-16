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

type MockManagerService struct{}

func (s *MockManagerService) CreateTeam(name string, managerID uint) (*database.Team, error) {
	team := &database.Team{
		Id:        1,
		Name:      name,
		ManagerID: managerID,
	}

	return team, nil
}

func (s *MockManagerService) GetAllTeams(managerID uint) (*[]database.Team, error) {
	teams := []database.Team{
		{
			Id:        1,
			Name:      "Team 1",
			ManagerID: managerID,
		},
		{
			Id:        2,
			Name:      "Team 2",
			ManagerID: managerID,
		},
	}

	return &teams, nil
}

func (s *MockManagerService) GetTeam(teamID uint, managerID uint) (*database.Team, error) {
	team := &database.Team{
		Id:        teamID,
		Name:      "Team",
		ManagerID: managerID,
	}

	return team, nil
}

func (s *MockManagerService) UpdateTeam(teamID uint, managerID uint, name string) (*database.Team, error) {
	team := &database.Team{
		Id:        teamID,
		Name:      name,
		ManagerID: managerID,
	}

	return team, nil
}

func (s *MockManagerService) DeleteTeam(teamID uint, managerID uint) error {
	return nil
}

func setupRouter(service IManagerService) *gin.Engine {
	router := gin.Default()
	controller := NewManagerController(nil, nil, service)

	router.POST("/teams", controller.CreateTeam)
	router.GET("/teams", controller.GetAllTeams)
	router.GET("/teams/:teamId", controller.GetTeam)
	router.PUT("/teams/:teamId", controller.UpdateTeam)
	router.DELETE("/teams/:teamId", controller.DeleteTeam)

	return router
}

func TestCreateTeam(t *testing.T) {
	service := &MockManagerService{}
	router := setupRouter(service)

	w := httptest.NewRecorder()
	body := BodyCreateTeam{
		Name: "Test Team",
	}
	bodyBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/teams?managerId=1", bytes.NewReader(bodyBytes))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseTeam database.Team
	err := json.Unmarshal(w.Body.Bytes(), &responseTeam)
	assert.NoError(t, err)
	assert.Equal(t, "Test Team", responseTeam.Name)
}

func TestGetAllTeams(t *testing.T) {
	service := &MockManagerService{}
	router := setupRouter(service)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/teams?managerId=1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseTeams []database.Team
	err := json.Unmarshal(w.Body.Bytes(), &responseTeams)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(responseTeams))
}

func TestGetTeam(t *testing.T) {
	service := &MockManagerService{}
	router := setupRouter(service)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/teams/1?managerId=1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var responseTeam database.Team
	err := json.Unmarshal(w.Body.Bytes(), &responseTeam)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), responseTeam.Id)
	assert.Equal(t, "Team", responseTeam.Name)
}
func TestUpdateTeam(t *testing.T) {
	service := &MockManagerService{}
	router := setupRouter(service)
	w := httptest.NewRecorder()
	body := BodyUpdateTeam{
		Name: "Updated Team"}
	bodyBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest("PUT", "/teams/1?managerId=1", bytes.NewReader(bodyBytes))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseTeam database.Team
	err := json.Unmarshal(w.Body.Bytes(), &responseTeam)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), responseTeam.Id)
	assert.Equal(t, "Updated Team", responseTeam.Name)
}

func TestDeleteTeam(t *testing.T) {
	service := &MockManagerService{}
	router := setupRouter(service)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/teams/1?managerId=1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Team deleted", response["message"])
}
