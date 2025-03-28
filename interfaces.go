package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/investment_calculator/todo"
)

// interface -> a contract that guarantees that have a certain method/s
// if interface has only one method - this is a name convention:
type Saver interface {
	Save() error
}

func SaveData(data Saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the data failed", err)
	} else {
		fmt.Println("Saving the data succeeded")
	}
	return err
}

type Displayer interface {
	Display()
}

// embedded interface
type Outputtable interface {
	Saver
	Displayer
}

func OutputData(data Outputtable) error {
	data.Display()
	err := SaveData(data)
	if err != nil {
		return err
	}
	return nil
}

// the special 'any value allowed' type
// value interface{}, alternatively instead you can also write any
// that's an alias for this special 'any value is allowed' type
func PrintSomething(value interface{}) {
	// type switch
	switch value.(type) {
	case int:
		// ...
		fmt.Println("Integer:", value)
	case string:
		fmt.Println("String:", value)
	default:
		fmt.Println("Other:", value)
	}

	// extracting type information from value
	typedVal, ok := value.(int)
	// type, boolean (it is okey) ^

	if ok {
		fmt.Println("It is integer", typedVal)
	}

	fmt.Println(value)
}

func Interfaces() {
	PrintSomething(0)
	PrintSomething(1)
	PrintSomething("Hello")
	content := getTodoInput("TODO: ")

	todo, err := todo.New(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	// todo.Display()
	// err = todo.Save()
	// if err != nil {
	// 	fmt.Println("Saving the todo failed", err)
	// } else {
	// 	fmt.Println("Saving the todo succeeded")
	// }

	// use generics
	// todo struct includes save method, so it works
	// it helps us to write more generic / usable code
	// err = SaveData(todo)
	// if err != nil {
	// 	return
	// }
	OutputData(todo)

	add(1, 2)
	add("A", "B")
	add(1.20, 2.4)
}

func getTodoInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	value := strings.TrimSuffix(text, "\n")
	value = strings.TrimSuffix(text, "\r")

	return value
}

// generics
func add[T string | int | float64](a, b T) T {
	return a + b
}
