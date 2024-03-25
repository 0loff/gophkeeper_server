package models

type UserLogin struct {
	UUID     string
	Email    string
	Password string
}

type UserAuth struct {
	ID       string
	Username string
	Password string
	Email    string
}
