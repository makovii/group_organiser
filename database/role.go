package database

type Role struct {
	Id      uint      `json:"id" gorm:"primary_key"`
	Role    string    `json:"role"`
}
