package dto

type Link struct {
	OriginalUrl string `json:"original_url" example:"www.google.com" binding:"required"`
	Slug        string `json:"slug" example:"abs1Ax" binding:"max=50"`
}

type UserLink struct {
	Id          int    `json:"id"`
	OriginalUrl string `json:"original_url" `
	Slug        string `json:"slug" `
}
