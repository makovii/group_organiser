package database

type Type struct {
	Id      uint      `json:"id" gorm:"primary_key"`
	Type    string    `json:"type"`
	Request []Request `json:"request`
}
