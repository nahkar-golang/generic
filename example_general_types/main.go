package general_types_example

import "fmt"

type Number interface {
	int64 | float64
}

type Numbers[T Number] []T

func ExampleWithGeneralTypes() {

	ints := Numbers[int64]{1, 2, 3}
	fmt.Println(ints)

	floats := Numbers[float64]{1.1, 2.2, 3.3}
	fmt.Println(floats)
}
