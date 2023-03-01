package database

type Status struct {
	Id         uint       `json:"id" gorm:"primary_key"`
	Status 		 string	    `json:"status"`
	Request  []Request    `json:"request`
}