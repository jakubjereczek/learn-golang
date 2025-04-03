package main

import "fmt"

func FunctionsDeepDive() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled, tripled) // [2 4 6 8] [3 6 9 12]

	doubled2 := transformNumbers(&numbers, getTransformerFunction())
	fmt.Println(doubled2)

	// anonymous functions (define a function just in time when you need it)
	doubled3 := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})
	fmt.Println(doubled3)

	doubled4 := transformNumbers(&numbers, createTransformer(2))
	fmt.Println(doubled4)

	// recursion
	fact := factorial(5)
	fmt.Println(fact)

	// variadic functions
	// sum := sumup(numbers)
	sum := sumup(1, 2, 3, 4, 5)
	fmt.Println(sum)

	// splitting slices into parameter values
	sum = sumup(numbers...)
	fmt.Println(sum)
}

type transformFn = func(int) int

// type anotherFn = func(int, []string, map[string][]int) ([]int, string)

// pass functions as params
func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

// returning functions as values
func getTransformerFunction() transformFn {
	return double
}

// factory function
func createTransformer(factor int) func(int) int {
	// closure here
	return func(number int) int {
		return number * factor
	}
}

// recursion
// factorial of 5 -> 5 * 4 * 3 * 2 * 1 => 120
func factorial(number int) int {
	// result := 1

	// for i := 1; i <= number; i++ {
	// 	result = result * i
	// }

	// return result

	if number == 0 {
		return 1
	}

	return number * factorial(number-1)
}

// variadic functions
// func sumup(numbers []int) int {
func sumup(numbers ...int) int { // merge it into a slice
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
