package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	Name string
	Age  int
}

/*
*
the usage of unsafe.Pointer is to convert a pointer to a uintptr,
then add an offset to the uintptr, and finally convert the uintptr back to a pointer.
*/
func main() {
	person := &Person{Name: "John", Age: 25}

	namePointer := unsafe.Pointer(&person.Name)
	agePointer := (*int)(unsafe.Pointer(uintptr(namePointer) + unsafe.Sizeof(person.Age)))

	*agePointer = 30

	fmt.Println(person.Age)  // Output: &{John 30}
	fmt.Println(person.Name) // Output: &{John 30}
}
