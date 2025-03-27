package user

import (
	"errors"
	"fmt"
	"time"
)

// to expose properties outside of struct, we need to UPPERCASE
// FirstName
// LastName
// ...
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// USER STRUCT METHOD! (user User) - 'Receiver'
func (user User) OutputUserDetails() {
	fmt.Println(user.firstName, user.lastName, user.birthDate)
}

// change data
// pass pointer!!!
func (user *User) ClearUserName() {
	// CHANGES DATA IN STRUCT
	user.firstName = ""
	user.lastName = ""
}

// USING CREATION / CONSTRUCTOR FUNCTIONS
// A UTILITY FUNCTION
// new.. is a name pattern
// func NewUser(name, lastName, birthDate string) User {
// 	return User{
// 		firstName: name,
// 		lastName:  lastName,
// 		birthDate: birthDate,
// 		createdAt: time.Now(),
// 	}
// }

// return a pointer (we prevent this value to be copied, it's not important for small piece of data like this but nice to know that feature) !!!
// IT IS OKEY TO JUST NAME THIS FUNCTION NEW (if it is in struct file!!!)
// func NewUser(name, lastName, birthDate string) (*User, error) {
func New(name, lastName, birthDate string) (*User, error) {
	// validation!!!
	if name == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("Invalid data")
	}

	return &User{
		firstName: name,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

// struct embedding
type Admin struct {
	email    string
	password string
	// or User User
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "firstName",
			lastName:  "lastName",
			birthDate: "01/01/1999",
			createdAt: time.Now(),
		},
	}
}
