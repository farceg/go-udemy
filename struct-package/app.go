package main

import (
	"fmt"

	"example.com/struct-package/user"
)

func main() {

	var appUser, err = user.New()

	if err == nil {
		fmt.Println("User:")
		appUser.OutputUserDetails()
		appUser.ClearUserName()
		fmt.Println("---------------------\nUser:")
		appUser.OutputUserDetails()
	} else {
		fmt.Println(err)
	}

}
