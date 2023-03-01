package database

type Player struct {
	Id         uint       `json:"id" gorm:"primary_key"`
	Name			 string     `json:"name"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	Request  []Request    `gorm:"foreignKey:From"`
}