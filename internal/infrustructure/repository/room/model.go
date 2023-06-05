package room

type Room struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Tag   string `json:"tag"`
	Owner User   `json:"owner"`
	Users []User `json:"users"`
}

type User struct {
	Code     string `json:"code"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
}
