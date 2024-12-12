package dto

type UsersLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignup struct {
	UsersLogin
	Phone string `json:"phone"`
}
