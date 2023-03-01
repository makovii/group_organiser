package database

type Request struct {
	Id         uint       `json:"id" gorm:"primary_key"`
	From			 uint       `json:"-"`
	To      	 uint       `json:"-"`
	StatusId   uint       `json:"-"`
	TypeId 		 uint       `json:"-"`
}