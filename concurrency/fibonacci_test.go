package concurrency

import "fmt"

func ExampleFibonacci() {
	for fib := range Fibonacci(15) {
		fmt.Print(fib, ", ")
	}
	// Output: 0, 1, 1, 2, 3, 5, 8, 13,
}
