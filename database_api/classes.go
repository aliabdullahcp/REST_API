package database_api

type Users struct {
	Id         int    `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
}

type UserRoles struct {
	Id        int    `json:"id"`
	User_Id   int    `json:"user_id"`
	User_Role string `json:"user_role"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)
