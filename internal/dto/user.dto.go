package dto

type User struct {
	Id            int    `json:"id"`
	Email         string `json:"email"`
	Fullname      string `json:"full_name"`
	Password_hash string `json:"password_hash"`
}
