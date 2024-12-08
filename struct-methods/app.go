package main

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

// Contructor for structs
func newUser() (*user, error) {

	inputFirstName := getUserData("Please enter your first name: ")
	inputLastName := getUserData("Please enter your last name: ")
	inputBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	if inputFirstName == "" || inputLastName == "" || inputBirthDate == "" {
		return nil, errors.New("No values passed")
	}

	return &user{
		firstName: inputFirstName,
		lastName:  inputLastName,
		birthDate: inputBirthDate,
		createdAt: time.Now(),
	}, nil
}

// Function to be used only by user struct
func (appUser user) outputUserDetails() {
	fmt.Println(appUser.firstName, appUser.lastName, appUser.birthDate)
}

// For editing the values it must receive a pointer
func (appUser *user) clearUserName() {
	appUser.firstName = ""
	appUser.lastName = ""
}

func main() {

	var appUser, err = newUser()

	if err == nil {
		fmt.Println("User:")
		appUser.outputUserDetails()
		appUser.clearUserName()
		fmt.Println("---------------------\nUser:")
		appUser.outputUserDetails()
	} else {
		fmt.Println(err)
	}

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
