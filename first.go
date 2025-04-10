package main // you need a package

// you need an entry point, because we are creating apps
func hello()  {
	println("Hello from Go",
	"Hello World!")
}

func init() {
	hello()
}


