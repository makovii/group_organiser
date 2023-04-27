package database

type User struct {
	Id       		uint      `json:"id" gorm:"primary_key"`
	Name     		string    `json:"name"`
	Email   		string    `json:"email"`
	Password 		string    `json:"password"`
	Request  		[]Request `gorm:"foreignKey:From"`
	Notification []string	`gorm:"type:text[]"`
	Teams				[]Team		`gorm:"foreignKey:Id"`
	Ban 				bool   		`json:"ban"`
	Role				int64    	`json:"role"`
}
