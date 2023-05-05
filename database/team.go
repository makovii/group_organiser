package database

type Team struct {
	Id           uint   `json:"id" gorm:"primary_key"`
	Name         string `json:"name"`
	ManagerID    uint   `json:"manager_id"`
	MemberIDs    []uint `json:"member_ids" gorm:"type:text[]"`
	UserRequests []uint `json:"user_requests" gorm:"type:text[]"`
}
