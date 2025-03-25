package main

import "fmt"

func Pointers() {
	// POINTER
	// What are Pointers?
	// variables that store value ADDRESSES instead of values

	age := 32
	// computer memory
	// value | address
	// 32    | 0xc00018050

	agePointer2 := &age
	fmt.Println("agePointer", agePointer2)

	// Why does this feature exist?
	// - it allows you to avoid unnecessary value copies (by default if you passed as an argument func a value it creates a copied value)
	// - it allows to directly mutate values (pass a pointer (address) instead of a value to a function - the function can then directly edit the underlying value - no return value is required)

	// e.g. without pointers (it copy value)
	ageY := 32
	adultYears := GetAdultYears(ageY)
	fmt.Println("Adult years", adultYears, ageY) // 14 32

	// // How can you work with Pointers?
	// e.g. with pointers
	agePointer := &ageY // *int

	// to get a value behind a pointer we need to use * before the pointer variable
	adultYearsWithPointer := GetAdultYearsPointer(agePointer)
	fmt.Println("Adult years (pointer)", adultYearsWithPointer, ageY) // 14 14

}

func GetAdultYears(age int) int {
	return age - 18
}

func GetAdultYearsPointer(age *int) int {
	// using pointers for data mutation
	*age -= 18

	return *age

	// return *age - 18
}

// a pointer's null value
// all values in go have a so-called "null value"
// for a pointer, it's nil - a special value build-into go
// nil represents the absence of an address - i.e., a pointer pointing at no address / no value in memory
