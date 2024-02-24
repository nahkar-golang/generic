package interface_example

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type CustomData interface {
	constraints.Ordered | []byte
}

type User[T CustomData] struct {
	ID   int
	Name string
	Data T
}

func ExampleWithInterface(name string) {
	u := User[int]{ID: 1, Name: name, Data: 10}
	u2 := User[string]{ID: 1, Name: name, Data: "10"}

	fmt.Println(u)
	fmt.Println(u2)
}
