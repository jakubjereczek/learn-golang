package main

import (
	"fmt"

	"example.com/investment_calculator/user"
)

func Structs() {
	firstName123 := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!

	// user := User{} // a null user/empty values

	/*
		var appUser User
		// it is called a struct literal
		appUser = User{
			firstName: firstName123,
			lastName:  lastName,
			birthDate: birthDate,
			createdAt: time.Now(),
		}
		// shortcut (remember to have correct order)
		// you can omit values, the type of this value will be nil
		appUser = User{
			firstName123,
			lastName,
			birthDate,
			time.Now(),
		}
	*/

	var appUser2 *user.User
	appUser2, err := user.New(firstName123, lastName, birthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test", "test")
	// admin.User.OutputUserDetails()
	// we can just simple use because the methods and properties are directly on the struct that embedded the struct
	admin.OutputUserDetails()

	// outputUserDetails(appUser)
	// outputUserDetails(&appUser) // pointer
	appUser2.OutputUserDetails() // struct method
	appUser2.ClearUserName()
	appUser2.OutputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

// func outputUserDetails(user User) {
// 	// ...
// 	fmt.Println(user.firstName, user.lastName, user.birthDate)
// }

// func outputUserDetails(user *User) {
// 	// ...

// 	// shortcut, it allowed by go (pointers to struct have it)
// 	// technically it is (*user).firstName
// 	fmt.Println((*user).firstName, user.lastName, user.birthDate)
// }
