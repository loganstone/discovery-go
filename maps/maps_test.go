package maps

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestCount1(t *testing.T) {
	codeCount := map[rune]int{}
	Count("가나다나", codeCount)

	if !reflect.DeepEqual(map[rune]int{'가': 1, '나': 2, '다': 1}, codeCount) {
		t.Error("Code count mismatch:", codeCount)
	}
}

func TestCount2(t *testing.T) {
	codeCount := map[rune]int{}
	Count("가나다나", codeCount)

	if len(codeCount) != 3 {
		t.Error("Code count ", codeCount)
		t.Fatal("count should be 3 but:", len(codeCount))
	}

	if codeCount['가'] != 1 || codeCount['나'] != 2 ||
		codeCount['다'] != 1 {
		t.Error("Code count mismatch:", codeCount)
	}
}

func ExampleCount1() {
	codeCount := map[rune]int{}
	Count("가나다나", codeCount)
	for _, key := range []rune{'가', '나', '다'} {
		fmt.Println(string(key), codeCount[key])
	}
	// Output:
	// 가 1
	// 나 2
	// 다 1
}

func ExampleCount2() {
	codeCount := map[rune]int{}
	Count("가나다나", codeCount)
	var keys sort.IntSlice
	for key := range codeCount {
		keys = append(keys, int(key))
	}
	sort.Sort(keys)
	for _, key := range keys {
		fmt.Println(string(key), codeCount[rune(key)])
	}
	// Output:
	// 가 1
	// 나 2
	// 다 1
}

func ExampleHasDupeRune() {
	fmt.Println(HasDupeRune("숨박꼭질"))
	fmt.Println(HasDupeRune("다시합시다"))
	// Output:
	// false
	// true
}
