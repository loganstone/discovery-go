package funcs

import (
	"fmt"
	"strings"
)

func ExampleReadFromToFunc() {
	r := strings.NewReader("bill\ntom\njane\n")
	f := func(line string) {
		fmt.Println(line)
	}
	err := ReadFromToFunc(r, f)
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// bill
	// tom
	// jane
}

func ExampleReadFromToCloser() {
	r := strings.NewReader("bill\ntom\njane\n")
	var lines []string
	f := func(line string) {
		lines = append(lines, line)
	}
	err := ReadFromToFunc(r, f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
	// Output:
	// [bill tom jane]
}

func ExampleNewGenerator() {
	gen := NewIntGenerator()

	fmt.Println(gen(), gen(), gen(), gen(), gen())
	fmt.Println(gen(), gen(), gen(), gen(), gen())
	// Output:
	// 1 2 3 4 5
	// 6 7 8 9 10
}

func ExampleNewGeneratorMultiple() {
	gen1 := NewIntGenerator()
	gen2 := NewIntGenerator()

	fmt.Println(gen1(), gen1(), gen1())
	fmt.Println(gen2(), gen2(), gen2(), gen2(), gen2())
	fmt.Println(gen1(), gen1(), gen1())
	// Output:
	// 1 2 3
	// 1 2 3 4 5
	// 4 5 6
}
