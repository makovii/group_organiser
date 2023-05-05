package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/makovii/group_organiser/config"
	"github.com/makovii/group_organiser/controller"
	"github.com/makovii/group_organiser/database"
	"github.com/makovii/group_organiser/middleware"
)


type IUserRepository interface {
	GetUserById(id int) (*database.User, error)
	GetNotifications(to int) (*[]database.Request, error)
	GetTeamById(id uint) (*database.Team)
	CreateRequest(request database.Request) (uint, error)
	SaveRequest(request database.Request) (uint, error)
	GetRequestsByFrom(from int) (*[]database.Request, error)
	GetAllManagers() (*[]database.User, error)
	GetAllTeams() (*[]database.Team, error)
}

type UserService struct {
	userRepository IUserRepository
}

func NewUserService(userRepository IUserRepository) *UserService {
	return &UserService{userRepository}
}

func (u *UserService) GetUserById(id int) (*database.User, error) {
	return u.userRepository.GetUserById(id)
}

func (u *UserService) GetNotifications(to int) (*[]database.Request, error) {
	return u.userRepository.GetNotifications(to)
}

func (u *UserService) JoinTeam(c *gin.Context, CFG *config.Config, body controller.BodyJoinTeam) (*database.Request, error) {
	authedUser, _ := c.Get("authedUser")
	user := authedUser.(middleware.AuthedUser)

	team := u.userRepository.GetTeamById((body.TeamId))

	if team.Id == 0 {
		return nil, errors.New("Team not found")
	}

	var request database.Request
	request.From = uint(user.Id)
	request.To = team.ManagerID
	request.StatusId = uint(CFG.Status.WaitId)
	request.TypeId = uint(CFG.Type.JoinTeamId)

	result, err := u.userRepository.CreateRequest(request)
	if err != nil {
		return nil, errors.New("Can't save request in db")
	}
	request.Id = result

	return &request, nil
}

func (u *UserService) LeaveTeam(c *gin.Context, CFG *config.Config, body controller.BodyJoinTeam) (*database.Request, error) {
	authedUser, _ := c.Get("authedUser")
	user := authedUser.(middleware.AuthedUser)

	team := u.userRepository.GetTeamById((body.TeamId))

	if team.Id == 0 {
		return nil, errors.New("Team not found")
	}

	var request database.Request
	request.From = uint(user.Id)
	request.To = team.ManagerID
	request.StatusId = uint(CFG.Status.WaitId)
	request.TypeId = uint(CFG.Type.LeaveTeamId)

	result, err := u.userRepository.CreateRequest(request)
	if err != nil {
		return nil, errors.New("Can't save request in db")
	}
	request.Id = result

	return &request, nil
}

func (u *UserService) CancelRequest(c *gin.Context, CFG *config.Config, id int) (*database.Request, error) {
	authedUser, _ := c.Get("authedUser")
	user := authedUser.(middleware.AuthedUser)

	userRequests, err := u.userRepository.GetRequestsByFrom(int(user.Id))
	if err != nil {
		return nil, errors.New("You don't have any requests yet")
	}

	for _, n := range *userRequests {
		if id == int(n.Id) {
			n.StatusId = uint(CFG.Status.CancelId)
			_, err := u.userRepository.SaveRequest(n)
			if err != nil {
				return nil, errors.New("Smth goes wrong with save request")
			}
			return &n, nil
		}
	}

	return nil, errors.New("You don't have request with this id")
}

func (u *UserService) GetAllManagers() (*[]database.User, error){
	managers, err := u.userRepository.GetAllManagers()
	if err != nil {
		return nil, errors.New("Smth goes wrong with get managers")
	}

	return managers, nil
}

func (u *UserService) GetAllTeams() (*[]database.Team, error){
	teams, err := u.userRepository.GetAllTeams()
	if err != nil {
		return nil, errors.New("Smth goes wrong with get managers")
	}

	return teams, nil
}

