package data

import "fmt"

var Countries [10]string // nil values when created
var Slice []int
var Codes map[int]bool // Maps, for every number you get a bool


func init() {
	Countries[0] = "South Africa"
	Countries[1] = "Brazil"
	Countries[8] = "USA"

	qty := len(Countries) // can use len function with different data types
	fmt.Println(qty)
	fmt.Println("Countries saved")
	fmt.Println(Countries)
	fmt.Println(Countries[1])

	Prices := [2]int { 100, 150 }
	fmt.Println(Prices)

	// Slices
	names := []string { "Mary", "John" }
	names = append(names, "Carol") //check this error, fixed by not using :=
	fmt.Println(names)

	wellKnownPorts := map[string]int {"http": 80, "https": 443} // JSON compatible
	fmt.Println(wellKnownPorts)
}