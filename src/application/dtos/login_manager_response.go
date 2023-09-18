package dtos


type LoginManagerResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}