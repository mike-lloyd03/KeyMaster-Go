package main

type Key struct {
	Name        string
	Description string
	Status      string
}

type User struct {
	ID           int
	Username     string
	Email        string
	DisplayName  string
	PasswordHash string
	CanLogin     bool
}
