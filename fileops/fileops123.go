package fileops

// to use other package name we need to create a folder with this name

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, filename string) {
	// convert to text
	valueText := fmt.Sprint(value)
	// convert to byte collection, permission like linux (owner read write)
	os.WriteFile(filename, []byte(valueText), 0644)
}

// returning a custom error
func GetFloatFromFile(filename string, defaultValue float64) (float64, error) {
	// errors
	// if error occurs it doesnt crash the application!
	// in go functions do not crash the application

	b, err := os.ReadFile(filename)
	// error handling:
	if err != nil {
		// if error is nil = no error

		// produce new error
		return defaultValue, errors.New("Failed to read file")
	}
	balanceText := string(b)
	res, _ := strconv.ParseFloat(balanceText, 64) // use _ to skip compilator check (not used...)

	return res, nil
}

// ONLY FUNCTIONS, VARIABLES that starts with UPPERCASE will be exposed (will be available in other packages)
