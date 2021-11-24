package main

import (
	"log"
)

type User struct {
	FullName     string
	Username     string
	PasswordHash string
}

type Storage interface {
	putOne(user *User) error
}

type App struct {
	Storage Storage
}

type SignUpType struct {
	Password        string
	ConfirmPassword string
	User
}

type AppError struct {
	Message string
}

func (err *AppError) Error() string {
	return err.Message
}

// func HashPassword(password string) (string, error) {

// }

func (app *App) SignUp(opts *SignUpType) error {
	if opts.Password != opts.ConfirmPassword {
		return &AppError{
			Message: "Passwords must match.",
		}
	}

	app.Storage.putOne(&opts.User)
	return nil
}

func main() {
	log.Println("Hello World")
}
