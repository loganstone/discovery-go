package main

import (
	"fmt"

	"github.com/hyunsuk/discovery-go/slice/eval"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	x := 9
	a = append(a, x)
	fmt.Println(a[3+1:])
	fmt.Println(a[3:])
	copy(a[3+1:], a[3:])
	a[3] = x

	fmt.Println(a)
	fmt.Println(eval.Eval("5"))
	fmt.Println("1 + 2")
	fmt.Println(eval.Eval("1 + 2"))
	fmt.Println("1 - 2 + 3")
	fmt.Println(eval.Eval("1 - 2 + 3"))
	fmt.Println("3 * ( 3 + 1 * 3 ) / 2")
	fmt.Println(eval.Eval("3 * ( 3 + 1 * 3 ) / 2"))
	fmt.Println("3 * ( ( 3 + 1 ) * 3 ) / 2")
	fmt.Println(eval.Eval("3 * ( ( 3 + 1 ) * 3 ) / 2"))
}
