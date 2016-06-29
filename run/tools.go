package run

import (
	"fmt"
	"time"

	"github.com/hyunsuk/discovery-go/slice/eval"
)

// CountDown is printing count down number.
func CountDown(seconds int) {
	for seconds > 0 {
		fmt.Println(seconds)
		time.Sleep(time.Second)
		seconds--
	}
}

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

	time.AfterFunc(3*time.Second, func() {
		fmt.Println("I'm so excited!!")
	})
	CountDown(5)
}
