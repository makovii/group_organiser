package database

type User struct {
	Id       		uint      `json:"id" gorm:"primary_key"`
	Name     		string    `json:"name"`
	Email   		string    `json:"email"`
	Password 		string    `json:"password"`
	Request  		[]Request `gorm:"foreignKey:From" json:"request"`
	Notification []string	`gorm:"type:text[]" json:"notification"`
	Teams				[]Team		`gorm:"foreignKey:Id" json:"teams"`
	Ban 				bool   		`json:"ban"`
	Role				int64    	`json:"role"`
}
