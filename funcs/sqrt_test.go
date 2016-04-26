package funcs

import "fmt"

func ExampleSqrt() {
	fmt.Printf("%.1f\n", Sqrt(2))
	fmt.Printf("%.1f\n", Sqrt(3))
	fmt.Printf("%.1f\n", Sqrt(4))
	// Output:
	// 1.4
	// 1.7
	// 2.0
}
