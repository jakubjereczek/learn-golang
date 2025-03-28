package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Printf("Todo task: \n%v", todo.Text)
}

func (note Todo) Save() error {
	fileName := "todo.json"
	jsonByteCollection, err := json.Marshal(note)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonByteCollection, 0644)
}

func New(text string) (*Todo, error) {
	if text == "" {
		return nil, errors.New("Empty input")
	}

	return &Todo{
		Text: text,
	}, nil
}
