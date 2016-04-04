package eval

import (
	"strconv"
	"strings"
)

// Eval retunrs the evaluation result of the given expr.
// The expression can have +, -, *, /, (, ) operators and
// decimal integers. Operators and operands should be
// space(" ") delimited.
func Eval(expr string) int {
	var ops []string
	var nums []int

	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}

	calculate := func(operator string) {
		b, a := pop(), pop()
		switch operator {
		case "+":
			nums = append(nums, a+b)
		case "-":
			nums = append(nums, a-b)
		case "*":
			nums = append(nums, a*b)
		case "/":
			nums = append(nums, a/b)
		}
	}

	reduce := func(higher string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]
			if strings.Index(higher, op) < 0 {
				// 목록에 없는 연산자 - 종료
				return
			}
			ops = ops[:len(ops)-1]

			if op == "(" {
				// 괄호 제거
				return
			}
			calculate(op)
		}
	}

	// for _, token := range strings.Split(expr, " ") {
	for _, s := range expr {
		if s == 32 {
			continue
		}
		token := string(s)
		switch token {
		case "(":
			ops = append(ops, token)
		case "+", "-":
			// 덧셈 & 뺄셈 이상의 우선순위를 가진 사칙연산 적용
			reduce("+-*/")
			ops = append(ops, token)
		case "*", "/":
			reduce("*/")
			ops = append(ops, token)
		case ")":
			reduce("+-*/(")
		default:
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
		// fmt.Printf("ops : %s\n", ops)
		// fmt.Printf("nums : %d\n", nums)
	}
	reduce("+-*/")
	return nums[0]
}
