package database

import (
	"gorm.io/gorm"
)

type Team struct {
	Id           uint   `json:"id" gorm:"primary_key"`
	Name         string `json:"name"`
	ManagerID    uint   `json:"manager_id"`
	MembersIDs   []uint `json:"members_ids" gorm:"type:text[]"`
	UserRequests []uint `json:"user_requests" gorm:"type:text[]"`
}

func (t *Team) HasUserRequest(userID uint) bool {
	for _, id := range t.UserRequests {
		if id == userID {
			return true
		}
	}
	return false
}
func (t *Team) AcceptUserRequest(userID uint, db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	t.MembersIDs = append(t.MembersIDs, userID)
	for i, id := range t.UserRequests {
		if id == userID {
			t.UserRequests = append(t.UserRequests[:i], t.UserRequests[i+1:]...)
			break
		}
	}

	if err := tx.Save(t).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
