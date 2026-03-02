package dto

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Request struct {
	Id string `json:"id"`
}
