package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	// struct tags - used to asign metadata (we will use them for output file)
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content: \n%v", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.TrimSpace(fileName)
	fileName = strings.ToLower(fileName) + ".json"
	println("fileName", fileName)
	// function to convert data to json, Marshal only transform the data publicly available (starts with uppercase)
	jsonByteCollection, err := json.Marshal(note)

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, jsonByteCollection, 0644)
	// return nil;
}

func New(title, content string) (*Note, error) {
	if title == "" || content == "" {
		return nil, errors.New("Empty input")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
