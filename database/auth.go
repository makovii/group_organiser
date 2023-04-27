package database

type Authentication struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role 		 int64	`json:"role"`
}

type Token struct {
	Email       string `json:"email"`
	TokenString string `json:"token"`
	Role				int64	`json:"role"`
}