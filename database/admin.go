package database

type Admin struct {
	Id       		uint      `json:"id" gorm:"primary_key"`
	Name     		string    `json:"name"`
	Email   		string    `json:"email"`
	Password 		string    `json:"password"`
	Notification []string		`json:"notification"`
}