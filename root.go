package main // in golang we need to split code into packages to organize code, we can export and import packages across the files
// why "main"? it's special package name that tells that this package will be the main entry point of the application

// build in packages: https://pkg.go.dev/std

// go standard packages:
// https://pkg.go.dev/std

// you can also install third-party libraries
// e.g. https://pkg.go.dev/github.com/brianvoe/gofakeit/v7
// install go get ...
// e.g. go get github.com/brianvoe/gofakeit

// go get - install

import (
	"errors"
	"fmt"
	"math"
	"os"

	// import local package
	"example.com/investment_calculator/fileops"
	// third-parties
	"github.com/brianvoe/gofakeit"
)

// this also must be named main (go knows which code to execute)
func main() {
	// basics()
	// constantsAndInput()
	// profitCalculator()
	// formattingStrings()
	// functions()

	// Bank()
	// Bank2()
	// WritingToFiles()

	// exit and throw error in console. it looks more like a crash
	// panic("Cannot continue. I'm sorry.")

	// ProfitCalculator2()
	// DummyData()

	Pointers()
}

// exec:
// - go run root.go
// - go build (go mod init), ./first-app (we can exec it without go installed)

// go.mod
// it tells that app belongs to this module path

// basic types:
// int - a number without decimal places
// float64 - a number with decinam places
// string - a text value,
// bool - true false

// niche basic types:
// uint - a unsigned integer which means a strictly non-negative integer (e.g. 0, 10, 255)
// int32 - a 32-bit signed integer, which is an integer with a specific range from -2,147,483,648 to 2,147,483,647 (e.g., -1234567890, 1234567890)
// rune - an alias for int32, represents a Unicode code point (i.e., a single character), and is used when dealing with Unicode characters (e.g., 'a', 'ñ', '世')
// uinit32 - a 32-bit unsigned integer, an integer that can represent values from 0 to 4,294,967,295
// int64 - a 64-bit signed integer, an integer that can represent a larger range of values, specifically from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
// there are more types like int8 and uint8 which work in the same way...

// null values;
// null values for diff types:
// int -> 0
// float64 -> 0.0
// string => ""
// bool - false

// run
// go run .

func basics() {
	var investmentAmount float64 = 1000
	// var expectedReturnRate = 5.5
	// shortcut (should be used as much as possible to declare when the type should be inferred by go - happy with infer type)
	expectedReturnRate := 5.5
	var years float64 = 10

	// go is staticly type lang
	// var featureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	// var featureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	featureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	// multiple variables in row
	var a, b float64 = 1024, 10
	a2, b2 := 1024.0, 10.0
	fmt.Println("ab", a, b, a2, b2)

	fmt.Println("Investment value", featureValue)
}

func constantsAndInput() {
	// not able to change const value stored in container
	const inflationRate = 2.5

	// var investmentAmount, years, expectedReturnRate float64 = 1000, 10, 5.5
	var investmentAmount, years, expectedReturnRate float64 // will be assigned to its null value, for float64 that is 0.0

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// load input
	// we need to point a pointer
	// fmt.Scan(&investmentAmount, &years, &expectedReturnRate)

	featureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	featureRealValue := featureValue / math.Pow(1+inflationRate/100, inflationRate)
	fmt.Println("Investment value with inflation rate", featureRealValue)
}

func profitCalculator() {
	revenue := GetValue("Type your revenue: ")
	expenses := GetValue("Type your expenses: ")
	taxRate := GetValue("Type your taxRate: ")

	ebt, profit, ratio := CalcProfit(revenue, expenses, taxRate)

	fmt.Println("Profit calculator (earnings before tax, earning after tax, ratio", ebt, profit, ratio)

}

func GetValue(info string) float64 {
	var input float64
	fmt.Print(info)
	fmt.Scan(&input)

	return input
}

func CalcProfit(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}

func formattingStrings() {
	revenue := 1626.5456323456
	otherValue := 2111

	// fmt.Printf("The expected revenue: %.2f.\n. The other value %v.", revenue, otherValue)

	formattedRevenue := fmt.Sprintf("The expected revenue: %.2f.\n", revenue)
	formattedOther := fmt.Sprintf("The other value %v.", otherValue)

	fmt.Print(formattedRevenue, formattedOther)

	// with `` we can add as many lines as we want
	fmt.Printf(`
	The expected revenue: %.2f. 
	The other value %v.`,
		revenue, otherValue)
}

func functions() {
	OutputText("Hello world!:)")

	revenue, future := CalculateFutureValue(3000, 5.5, 5, 2.5)
	// revenue, future := CalculateFutureValue2(3000, 5.5, 5, 2.5)

	fmt.Printf(`The expected revenue: %.2f. 
	The future value %v.`, revenue, future)
}

func OutputText(payload string) {
	fmt.Print(payload)

}

func CalculateFutureValue(investimentAmount, expectedReturnRate, years, inflationRate float64) (float64, float64) {
	fv := investimentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)

	// we can return few values, use comma
	return fv, rfv
}

// alternative syntax for returns

func CalculateFutureValue2(investimentAmount, expectedReturnRate, years, inflationRate float64) (fv float64, rfv float64) {
	fv = investimentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	// return fv, rfv
	return // go automatically returns these two values
}

// we can declare constant and variables outside of the functions
const INFLATION_RATE = 2.5

// Control Structures - it allows to control which things are executed
func Bank() {
	var accountBalance float64 = 1000

	// for i := 0; i < 2; i++ {
	for {
		// import from other file
		// PresentOptions()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		fmt.Println("You choose", choice)

		if choice == 1 {
			fmt.Println("* Your balance is", accountBalance)
		} else if choice == 2 {
			fmt.Println("* How much do you want to deposit?")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("* Invalid amount")
				// return
				continue
			}

			accountBalance += depositAmount
			fmt.Println("* Balance updated. New amount:", accountBalance)
		} else if choice == 3 {
			fmt.Println("* How much do you want to withdraw?")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 || withdrawAmount > accountBalance {
				fmt.Println("* Invalid amount")
				// return
				continue // skip current iteration and continue in the next
			}

			accountBalance -= withdrawAmount
			fmt.Println("* Balance updated. New amount:", accountBalance)
		} else {
			fmt.Println("* Exit")
			// return
			break // break allows you to execute code below
		}
	}

	fmt.Println("Thanks for choosing our bank.")
}

// it moved into package
// func PresentOptions() {
// 	fmt.Println("Welcome to GO bank.")
// 	fmt.Println("What do you want to do?")
// 	fmt.Print("1. Check balance \n2. Deposit money \n3. Withdraw money \n4. Exit \n")
// }

// go has only for loop (there is no while construction)
// * for i := 0; i < 2; i++ { }
// * infinity loop - for {}
// * condition loop - for someCondition {}
// someCondition is an expression that yields a boolean value or variable that contains a boolean value

func Bank2() {
	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	// break is not necessary in go (for switch ofc)
	// if you a re using break it exit the calling code below the switch statement (out of the loop)
	switch choice {
	case 1:
		fmt.Println("Case 1")
		// break
	case 2:
		fmt.Println("Case 2")
		// break
	case 3:
		fmt.Println("Case 3")
		// break
	default:
		fmt.Println("Case X")
		// break
	}
}

func WritingToFiles() {
	// use func from local file
	balance, err := fileops.GetFloatFromFile(file, 1000)

	if err != nil {
		fmt.Println("Error")
		fmt.Println("We wil use the default balance", err)
	}

	fmt.Println("Current balance", balance)

	var amount float64
	fmt.Print("Change balance: ")
	fmt.Scan(&amount)

	fileops.WriteFloatToFile(amount, file)
}

const file = "balance.txt"

// more advanced

func ProfitCalculator2() {
	revenue, err := GetValue2("Type your revenue: ")
	expenses, err_2 := GetValue2("Type your expenses: ")
	taxRate, err_3 := GetValue2("Type your taxRate: ")

	if err != nil || err_2 != nil || err_3 != nil {
		fmt.Println(err, err_2, err_3)
		return
	}

	ebt, profit, ratio := CalcProfit(revenue, expenses, taxRate)

	fmt.Println("Profit calculator (earnings before tax, earning after tax, ratio", ebt, profit, ratio)

	StoreResults(ebt, profit, ratio)

}

func GetValue2(info string) (float64, error) {
	var input float64
	fmt.Print(info)
	fmt.Scan(&input)

	if input <= 0 {
		// default value, error
		return 0, errors.New("Not supported value")
	}

	return input, nil
}

func StoreResults(ebt, profit, ratio float64) {
	txt := fmt.Sprintf("EBT: %.1f\nPROFIT: %.f\nRATIO: %.3f\n", ebt, profit, ratio) // generate new string
	os.WriteFile("results.txt", []byte(txt), 0644)
}

func DummyData() {
	fmt.Println("Dummy CarModel", gofakeit.CarModel())
}
