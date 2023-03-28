package database

type Team struct {
	Id       	 uint   `json:"id" gorm:"primary_key"`
	Name       string `json:"name"`
	ManagerID  uint   `json:"manager_id"`
	MembersIDs []uint `json:"members_ids" gorm:"type:text[]"`
}