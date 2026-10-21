package dto

type User struct {
	Id            int    `json:"id"`
	Email         string `json:"email"`
	Password_hash string `json:"password_hash"`
}
