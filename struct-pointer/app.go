package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {

	var appUser = user{
		firstName: getUserData("Please enter your first name: "),
		lastName:  getUserData("Please enter your last name: "),
		birthDate: getUserData("Please enter your birthdate (MM/DD/YYYY): "),
		createdAt: time.Now(),
	}

	outputUserDetails(&appUser)

}

func outputUserDetails(appUser *user) {
	fmt.Println(appUser.firstName, appUser.lastName, appUser.birthDate, appUser.createdAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
