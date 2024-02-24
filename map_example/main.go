package map_example

import "fmt"

type CustomMap[T comparable, V int | string] map[T]V

func ExampleWithMap() {

	m := CustomMap[string, int]{"a": 10, "b": 20}
	fmt.Println(m)

	m1 := CustomMap[string, string]{"a": "10", "b": "20"}
	fmt.Println(m1)

	m2 := CustomMap[int, int]{1: 10, 2: 20}
	fmt.Println(m2)
}
