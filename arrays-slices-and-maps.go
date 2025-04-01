package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func Lists() {
	prices := [4]float64{10.99, 9.99, 44.95, 20.49}
	fmt.Println(prices) // [10.99 9.99 44.95 20.49]

	var productNames [3]string
	productNames = [3]string{"A Book"}
	fmt.Println(productNames) // [A Book  ]

	// Accessing by index
	fmt.Println(prices[2]) // 44.95

	// Set by index
	productNames[1] = "A Second Book"
	fmt.Println(productNames) // [A Book A Second Book ]

	// Selecting Parts of Arrays With Slices
	// Kind a new Array based on existing one
	featuredPrices := prices[1:3] // [9.99 44.95] (from one to three but excluding the last one)
	fmt.Println(featuredPrices)

	// take until the third (excluded)
	// prices[:3]
	fmt.Println(prices[:3]) // [10.99 9.99 44.95]

	// take from the index x to the end
	// prices[1:]
	fmt.Println(prices[1:]) // [9.99 44.95 20.49]

	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices) // [9.99]

	// Slices are like a reference, a bit like a pointer, a window into an array.
	// If we modify element in a slice, we would also modify element in ORIGINAL array.
	featuredPrices[0] = 2115.00
	fmt.Println(prices)            // [10.99 2115 44.95 20.49]
	fmt.Println(highlightedPrices) // [2115]

	// When you create a slice you do not copy the original array but just have one array in memory, the slice is just a reference.

	// length, capacity (you can always select more torward the end of an array, but not torwars the start - so torwards the left, capacity only counts torward the end of original array)
	fmt.Println(featuredPrices, len(featuredPrices), cap(featuredPrices)) // [2115, 44.95] 2 3

	// building dynamic list with slices
	// omit the length of array, go actually automatically create a slice for you, and since a slice always is based on an array, it will also create an array for you behind the scenes (but it will automatically ditch that array, and create a new array if your slice grows)
	dPrices := []float64{11.95, 4.95}
	fmt.Println(dPrices[0])
	// dPrices[2] = 10 .0// error: index of of range
	// instead of: use append fn that add an item to that slice and returns a new brand array
	updatedDPrices := append(dPrices, 10.0)
	fmt.Println(dPrices, updatedDPrices) // [11.95 4.95] [11.95 4.95 10]

	// remove first element from array (slice syntax)
	dPrices = prices[1:]

	// practice slices
	ListsPractice()

	// unpacking list values
	products := []float64{10.99, 4.95}
	products = append(products, 12.95, 29.95, 100.0)
	products = products[1:]
	discountPrices := []float64{1, 2, 3, 4}
	products = append(products, discountPrices...) // <- unpacking
	fmt.Println(products)                          // [4.95 12.95 29.95 100 1 2 3 4]

}

func ListsPractice() {
	//  1. Create a new array (!) that contains three hobbies you have
	//     Output (print) that array in the command line.
	hobbies := [3]string{"Gym", "Programming", "Football"}
	fmt.Println(hobbies)

	//  2. Also output more data about that array:
	//     - The first element (standalone)
	//     - The second and third element combined as a new list
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	//  3. Create a slice based on the first element that contains
	//     the first and second elements.
	//     Create that slice in two different ways (i.e. create two slices in the end)
	selectedHobbies := hobbies[0:2]
	selectedHobbies2 := hobbies[:2]
	fmt.Println(selectedHobbies, selectedHobbies2) // [Gym Programming] [Gym Programming]

	//  4. Re-slice the slice from (3) and change it to contain the second
	//     and last element of the original array.
	fmt.Println(selectedHobbies)      // [Gym Programming]
	fmt.Println(cap(selectedHobbies)) // 3

	resliceHobbies := selectedHobbies[1:3]
	fmt.Println(resliceHobbies) // [Programming Football]

	//  5. Create a "dynamic array" that contains your course goals (at least 2 goals)
	courseGoals := []string{"Learn Go"}
	courseGoals = append(courseGoals, "Create REST API in Go")
	fmt.Println(courseGoals) // [Learn Go Create REST API in Go]

	//  6. Set the second goal to a different one AND then add a third goal to that existing dynamic array
	courseGoals[1] = "Learn all the details"
	courseGoals = append(courseGoals, "Practice Go")
	fmt.Println(courseGoals) // [Learn Go Learn all the details Practice Go]

	//  7. Bonus: Create a "Product" struct with title, id, price and create a
	//     dynamic list of products (at least 2 products).
	//     Then add a third product to the existing list of products.
	products := []Product{
		// Product{id: "1", title: "product title", price: 10.99},
		// Product{id: "2", title: "product title", price: 19.99},
		{id: "1", title: "product title", price: 10.99},
		{id: "2", title: "product title", price: 19.99},
	}
	fmt.Println(products) // [{product title 1 10.99} {product title 2 19.99}]
	products = append(products, Product{
		id:    "3",
		title: "product title",
		price: 4.95,
	})
	fmt.Println(products) // [{product title 1 10.99} {product title 2 19.99} {product title 3 4.95}]
}

func Maps() {
	// A map is a different kind of data structure (use to group data together).
	// websites := []string{"https://www.udemy.com/course/go-the-complete-guide", "https://www.google.com"}
	websites := map[string]string{
		"Udemy Course": "https://www.udemy.com/course/go-the-complete-guide",
		"Google":       "https://www.google.com",
	}
	fmt.Println(websites) // map[Google:https://www.google.com Udemy Course:https://www.udemy.com/course/go-the-complete-guide]

	// accessing
	fmt.Println(websites["Udemy Course"])

	// map is dynamic like slices (we can always add new key value pairs)
	// any value can be used as a key

	// assigning
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites) // map[Google:https://www.google.com LinkedIn:https://linkedin.com Udemy Course:https://www.udemy.com/course/go-the-complete-guide]

	// deleting
	delete(websites, "Udemy Course")
	fmt.Println(websites) // map[Google:https://www.google.com LinkedIn:https://linkedin.com]

}

// using the special `make` function
// go creates a bigger array and then does not have to recreate those arrays all the time, which is a bit more efficient
// by default in the case of array append() creates a new array

func Make() {
	// 1. type of slice, 2. length of slice - initial, 3 - capacity (the capacity is essentially the maximum number of array items and therefore controls how much memory space allocate)
	userNames := make([]string, 2, 10)
	// userNames = append(userNames, "Jacoob")
	// userNames = append(userNames, "Max")
	userNames[0] = "Jacoob"
	userNames[1] = "Max"

	fmt.Println(userNames)

	// to sum up: this make function can be useful if you know in advance that you are soon going to add a fixed number of items or at least a number where you have a rough estimate how much you going to add
	// that can make memory management more efficient

	// maps
	// courseRatings := map[string]float64{}
	// the same problem as we have with arrays, when we create such an empty map here, go will have to reallocate memory whenever we add new items to that map
	// courseRatings := make(map[string]float64, 3)
	courseRatings := make(floatMap, 3)
	courseRatings["Go"] = 4.7
	courseRatings["CSS"] = 4.8
	courseRatings["Vue"] = 4.5
	// ...then if we exceeded the size number, go needs to reallocate the map
	courseRatings["React"] = 4.5

	courseRatings.output()

}

// type alias
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func Loops() {
	// for loops with arrays, slices and maps
	coordinates := map[string][2]float64{
		"Gdansk": {54.3722, 18.6386},
		"Gdynia": {54.5189, 18.5305},
		"Sopot":  {54.4418, 18.5601},
	}

	// if you do not care about the individual item values or indexes you can also jus write
	// for range coordinates {}
	for key, value := range coordinates {
		fmt.Println(key, value) // Gdansk [54.3722 18.6386]...
	}

	// arrays / slices
	hobbies := [3]string{"Gym", "Programming", "Football"}

	for index, value := range hobbies {
		fmt.Println(index, value) // 0 Gym...
	}
}
