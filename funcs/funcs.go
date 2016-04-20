package funcs

import (
	"bufio"
	"io"
)

// ReadFrom is Read some string slice
func ReadFrom(r io.Reader, lines *[]string) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		*lines = append(*lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// ReadFromToFunc is use some Func
func ReadFromToFunc(r io.Reader, f func(lines string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// NewIntGenerator is return generator number
func NewIntGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}
