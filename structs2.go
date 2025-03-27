package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/investment_calculator/note"
)

// creating other custom types & adding methods

// alias
type customString string

// function for our type
func (text customString) log() { // it wont be possible with built-in type
	fmt.Println(text)
}

func Structs2() {
	var name customString = "jakub"
	name.log()

	// practice
	StructsPractice()
}

func StructsPractice() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	notee := *userNote

	if err != nil {
		fmt.Println(err)
		return
	}

	notee.Display()
	err = notee.Save()
	if err != nil {
		fmt.Println("Saving the note failed", err)
	} else {
		fmt.Println("Saving the note succeeded")
	}
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	// var value string
	// fmt.Scanln(&value)

	// alternative package for user input
	// os.Stdin - standard command line
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n') // read until \n, use -> ''
	if err != nil {
		return ""
	}
	value := strings.TrimSuffix(text, "\n")
	value = strings.TrimSuffix(text, "\r")

	return value
}
