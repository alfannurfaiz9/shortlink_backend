package dto

type Register struct {
	Fullname string `json:"full_name" example:"your name" binding:"required"`
	Email    string `json:"email" example:"yourname@mail.com" binding:"required,email"`
	Password string `json:"password" example:"password" binding:"required,min=6"`
}

type Login struct {
	Email    string `json:"email" example:"yourname@mail.com"`
	Password string `json:"password" example:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UserInfo struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}
