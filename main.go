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

func calculateTaxWithName(price float32) (stateTax float32, cityTax float32) {
	stateTax = price*0.9
	cityTax = price*0.02
	return
}

//pointers
func increment(x *int) {
	*x++
}

func birthday(pointerAge *int) { // pointer, denoted with an asterisk
	// print(pointerAge)
	// print(*pointerAge)

	if (*pointerAge>140) {
		// panic goes up the callstack and there are ways to trap that panic
		// closes app with a message, this is not how you do error management
		panic("Too old to be true") // it is an inturruption aborting, like an exception



	}

	//string format
	fmt.Printf("The pointer is %v and the value is %v\n", pointerAge, *pointerAge) // expecting int: %i (error), string: %s, any: %v



	*pointerAge = *pointerAge + 1
	//age ++ // does not work

}


// you need an entry point, because we are creating apps
func main()  {

	// it is like LIFO
	defer fmt.Println("Good ")
	defer fmt.Println("Bye!!") //executes at the end, if defers the execution until the end of the function call is reached

	/*
	defer is best used when you will forget to do something, e.g opening and closing the database
	typically you forget so you open it and (defer) close it immediately and defer the closing

	defer works with panic(), if there is a panic it will execute your defers before closing the app
	it honors your defers
	*/


	stateTax, cityTax := calculateTax(100) // create two variables on the fly
	tax1, _ := calculateTax(200) // underscore means you don't need it, ignore it
	fmt.Println(stateTax, cityTax)
	fmt.Println(tax1)

	fmt.Println(addAndSubtract(2, 3))
	fmt.Println(add(1,2))
	println("Hello from a module")

	fmt.Println(calculateTaxWithName(12))

	//i := 1
	//fmt.Println(increment(i))

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

	// Collections, arrays, in collections.go
	// valnum := 0
	var valnum int = 0
	increment(&valnum)
	fmt.Println(valnum)


	//Pointers
	age := 22
	birthday(&age) // prefix of & for pointers, cloning the value, passing it as a copy, incrementing the copy not the variable. Getting the adress of the pointer
	fmt.Println(age)
	fmt.Println(&age) // memory adress of the age

	//fmt.Printf()

	

}