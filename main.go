package main

import (
	"fmt"

	"./examples"
)

func main() {
	// fmt.Print("Hello World\n")
	// testFor(10)
	// fmt.Print(testArray())
	// testSlice()
	// m, err := testMap(1, 2, 3, 4)
	// fmt.Print(m, err)
	// testStruct()
	// testItf()
	// testGoroutine()
	fmt.Printf("Main Func\n")
	examples.TestChannel()
}
