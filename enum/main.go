package main

import "fmt"

const (
	Red             = (1 << iota)
	Green           = (1 << iota)
	Blue, ColorMask = (1 << iota), (1 << (iota + 1)) - 1
)

const (
	i complex128 = complex(1000, 1000)
)

func main() {

	fmt.Println("Red:", Red)
	fmt.Println("Green:", Green)
	fmt.Println("Blue:", Blue)
	fmt.Println("ColorMask:", ColorMask)

	fmt.Println("i:", i)
}

// Output:
//Red: 1
//Green: 2
//Blue: 4
//ColorMask: 7
//i: (1000+1000i)
