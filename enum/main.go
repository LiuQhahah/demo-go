package main

import "fmt"

const (
	Red             = (1 << iota)
	Green           = (1 << iota)
	Blue, ColorMask = (1 << iota), (1 << (iota + 1)) - 1
)

func main() {

	fmt.Println("Red:", Red)
	fmt.Println("Green:", Green)
	fmt.Println("Blue:", Blue)
	fmt.Println("ColorMask:", ColorMask)
}

// Output:
//Red: 1
//Green: 2
//Blue: 4
//ColorMask: 7
