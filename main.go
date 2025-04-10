package main // you need a package

import "fmt" // format package

import "frontendmasters.com/go/io/data" // need to add prefix of project id, takes last part of URL as name of the object to reference things there, or can use an alias

var name = "Frontend Masters" // not a global variable but a package variable

var url = "https://frontendmasters.com" //global variable, in a package

func init() {
	fmt.Println("A")
}

func init() {
	fmt.Println("B")
}

func add(a int, b int) int {
	return a+b
}

func addAndSubtract(a int, b int) (int, int) {
	return a+b, a-b
}

func calculateTax(price float32) (float32, float32) {
	return price*0.9, price*0.02
}

// you need an entry point, because we are creating apps
func main()  {
	stateTax, cityTax := calculateTax(100) // create two variables on the fly
	tax1, _ := calculateTax(200) // underscore means you don't need it, ignore it
	fmt.Println(stateTax, cityTax)
	fmt.Println(tax1)

	println("Hello from a module")

	printData()

	println(name)


	var x int //variable declaration
	var name string
	const y = 2

	message := "Hello from Go"
	price := 34.4

	const pi = 3.14

	var z int = 2

	var text string = "Hello!"

	//const maxSpeed byte = 60

	otherText := "Bye!" // variable initialization shorcut, no var or type, makes about two lines shorter
	// only works inside functions

	isTrue := true
	var isFalse bool = false

	print(isFalse && isTrue)
	print(isFalse || isTrue)
	print(!isFalse)
	print(y >= z)

	fmt.Print(data.MaxSpeed)

	print(message, price, x, name, z, text, otherText, pi, /*maxSpeed,*/ url, isTrue, isFalse)


	// no automatic conversion between types
	id := 4

	priceof := 45.4

	priceAsInt := int(priceof)
	idAsFloat := float32(id)

	fmt.Println()
	fmt.Println()
	fmt.Println(priceAsInt)
	fmt.Println(idAsFloat)

	str1 := "This is just a text"

	str2 := `Every text
	is multiline
	by default`

	fmt.Println(str1)

	fmt.Println(str2)

	fmt.Print()
	data.ShowOutputFunc() // from data package

	// Collections, arrays

	

}