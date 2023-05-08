package repository

import (
	"errors"

	"github.com/makovii/group_organiser/config"
	"github.com/makovii/group_organiser/database"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
	CFG *config.Config
}

func NewUserRepository(db *gorm.DB, cfg *config.Config) *UserRepository {
	return &UserRepository{DB: db, CFG: cfg}
}

func (u *UserRepository) GetUserById(id int) (*database.User, error) {
	var user database.User
	if err := u.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return &user, err
	}
	return &user, nil
}

func (u *UserRepository) GetNotifications(to int) (*[]database.Request, error) {
	var requests []database.Request
	if err := u.DB.Where("\"to\" = ?", to).Find(&requests).Error; err != nil {
		return &requests, err
	}
	return &requests, nil
}

func (u *UserRepository) GetTeamById(id uint) (*database.Team) {
	var team database.Team
	u.DB.Where("id = ?", id).First(&team)
	return &team
}

func (u *UserRepository) CreateRequest(request database.Request) (uint, error) {
	requestToDB := u.DB.Create(&request)

	if requestToDB.Error != nil {
		return 0, errors.New("Can't create request in db")
	}

	return request.Id, nil
}

func (u *UserRepository) SaveRequest(request database.Request) (uint, error) {
	requestToDB := u.DB.Save(&request)

	if requestToDB.Error != nil {
		return 0, errors.New("Can't save request in db")
	}

	return request.Id, nil
}

func (u *UserRepository) GetRequestsByFrom(from int) (*[]database.Request, error) {
	var requests []database.Request
	if err := u.DB.Where("\"from\" = ?", from).Find(&requests).Error; err != nil {
		return &requests, err
	}
	return &requests, nil
}

func (u *UserRepository) GetAllManagers() (*[]database.User, error) {
	var managers []database.User
	if err := u.DB.Where("role = ?", u.CFG.Role.ManagerId).Find(&managers).Error; err != nil {
		return &managers, err
	}
	return &managers, nil
}

func (u *UserRepository) GetAllTeams() (*[]database.Team, error) {
	var teams []database.Team
	if err := u.DB.Find(&teams).Error; err != nil {
		return &teams, err
	}
	return &teams, nil
}





