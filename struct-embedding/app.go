package main

import (
	"fmt"

	"example.com/struct-package/user"
)

func main() {

	var admin, _ = user.NewAdmin("Felipe", "Arce")
	admin.OutputUserDetails()

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
