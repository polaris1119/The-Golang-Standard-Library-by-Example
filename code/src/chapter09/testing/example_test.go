package testing_test

import (
	"fmt"

	. "../testing"
)

func ExampleFib() {
	fmt.Println(Fib(7))
	// Output: 13
}
