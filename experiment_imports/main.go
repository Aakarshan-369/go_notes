package main

import (
	"experiment_imports/utils"
	"fmt"
)

// Go is a compiled, statically typed language with a C­like syntax and garbage collection.
// Go does not need a runtime, it compiles directly to machine code
// Go code is organized into packages, which are collections of source files in the same directory
// the go.mod file is used to manage dependencies
// the module is the collection of packages that are used to build a program
// a single module typically contains multiple packages

func main() {
	// directly call the SayHello function because its the same package "main"
	// and it is in the same directory
	// Go's import system works at the package level, not the file level
	// They're in the same directory, allowing direct function calls between files
	message := functionThatCanBeCalledinMain() // accessible in main.go without explicit import since they share the same package
	fmt.Println(message)
	fmt.Println("Hello from main!")

	// an import needed from a different package in a different directory
	// this import function name shoudl have Capital Start
	// now this whole package main can be compiled and run using `go run .`
	// if we do `go run main.go` that creates only compiles the main file (can't fint the function in hello.go)
	messageFromUtils := utils.CapitalStartingFunctionIsFromUtilsPackage()
	fmt.Println(messageFromUtils)
}
