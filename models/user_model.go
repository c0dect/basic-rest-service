package models

type User struct {
	UserId   int64  `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
}
