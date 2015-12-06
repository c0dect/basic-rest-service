package models

type Errors struct {
	Errors []*Error `json:"errors"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}
