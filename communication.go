// every go file must be a part of a package!

// you be able to call it in every file that use main package
package main

import "fmt"

func PresentOptions() {
	fmt.Println("Welcome to GO bank.")
	fmt.Println("What do you want to do?")
	fmt.Print("1. Check balance \n2. Deposit money \n3. Withdraw money \n4. Exit \n")
}

// package main
// package utility
