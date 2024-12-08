package user

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// Name capitalized to access the struct from other files and packages
type Admin struct {
	email    string
	password string
	user
}

// To expose the atributes they muste be capitalized
func NewAdmin(email, password string) (*Admin, error) {
	return &Admin{
		email:    email,
		password: password,
		user: user{
			firstName: "Admin",
			lastName:  "Admin",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}, nil
}

// Contructor for structs
func New() (*user, error) {

	inputFirstName := getUserData("Please enter your first name: ")
	inputLastName := getUserData("Please enter your last name: ")
	inputBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	if inputFirstName == "" || inputLastName == "" || inputBirthDate == "" {
		return nil, errors.New("no values passed")
	}

	return &user{
		firstName: inputFirstName,
		lastName:  inputLastName,
		birthDate: inputBirthDate,
		createdAt: time.Now(),
	}, nil
}

// Function to be used only by user struct
func (appUser user) OutputUserDetails() {
	fmt.Println(appUser.firstName, appUser.lastName, appUser.birthDate)
}

// For editing the values it must receive a pointer
func (appUser *user) ClearUserName() {
	appUser.firstName = ""
	appUser.lastName = ""
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
