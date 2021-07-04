package models

type Person struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
}