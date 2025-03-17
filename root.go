package main // in golang we need to split code into packages to organize code, we can export and import packages across the files
// why "main"? it's special package name that tells that this package will be the main entry point of the application

// build in packages: https://pkg.go.dev/std

import "fmt"

// this also must be named main (go knows which code to execute)
func main() {
	fmt.Print("Hello world") // common choice
	fmt.Print(`Hello world`)
}

// exec:
// - go run root.go
// - go build (go mod init), ./first-app (we can exec it without go installed)

// go.mod
// it tells that app belongs to this module path
