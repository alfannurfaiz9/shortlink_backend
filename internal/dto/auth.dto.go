package dto

type Register struct {
	Fullname string `json:"full_name" example:"your name" binding:"required"`
	Email    string `json:"email" example:"yourname@mail.com" binding:"required,email"`
	Password string `json:"password" example:"password" binding:"required,min=6"`
}
