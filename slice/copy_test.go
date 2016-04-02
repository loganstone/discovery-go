package slice

import "fmt"

func Example_copy() {
	src := []int{1, 2, 3, 4, 5, 6}
	dest := make([]int, len(src))

	if n := copy(dest, src); n != len(src) {
		fmt.Println("복사가 덜됐습니다")
	}

	fmt.Println(dest)
	// Output:
	// [1 2 3 4 5 6]
}
